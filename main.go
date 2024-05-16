package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/joho/godotenv"
)

const input_csv = "mensajes_analizar.csv";
const num_workers = 20

func main() {

	fmt.Printf("\nRecuerda tener un fichero \"%s\" en el mismo directorio que este programa. Si no, la ejecución fallará\n\n", input_csv)

	output_csv := ""
	num_existing_lines := 0
	scanner := bufio.NewScanner(os.Stdin)
	
	for { 
		fmt.Print("Introduce el nombre del fichero de salida.\nSi ya has analizado líneas del CSV de entrada, utiliza el mismo nombre de salida para comenzar donde terminó: ")
		scanner.Scan()
		output_csv = strings.Trim(scanner.Text(), " ")
		if output_csv == "" {
			fmt.Print("Introduce un nombre válido, sin espacios ni barras.\n")
			continue
		}

		num_existing_lines = get_num_existing_lines(output_csv)
		if num_existing_lines == -1 {
			fmt.Print("El fichero de salida no existe. Esto analizará tu CSV desde el inicio. ¿Quieres continuar? (y/n): ")
			scanner.Scan()
			resp_input := scanner.Text()
			for resp_input != "y" && resp_input != "n" {
				fmt.Print("¿Quieres continuar? (y/n): ")
				scanner.Scan()
				resp_input = scanner.Text()
			}

			if resp_input == "y" {
				break
			}
		} else {
			fmt.Printf("CSV de salida con %d líneas ya analizadas\n", num_existing_lines)
			fmt.Print("El programa comenzará a analizar el CSV desde la línea siguiente. ¿Quieres continuar? (y/n): ")
			scanner.Scan()
			resp_input := scanner.Text()
			for resp_input != "y" && resp_input != "n" {
				fmt.Print("¿Quieres continuar? (y/n): ")
				scanner.Scan()
				resp_input = scanner.Text()
			}

			if resp_input == "y" {
				break
			}
		}
	}


	// Open the CSV file
	file, err := os.Open(input_csv)
	if err != nil {
		log.Fatalf("Error abriendo csv: %s. ¿Está en el mismo directorio que el ejecutable?", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Open the existing CSV file for appending
	outputFile, err := os.OpenFile(output_csv, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error abriendo csv de salida %s: %s.", output_csv, err)
	}
	defer outputFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Order dictionary to find longer emotes first
	// Example: rubbComo and rubbC collide, so we need to check first the longest
	sort.Slice(emotes_dictionary, func(i, j int) bool {
		return len(emotes_dictionary[i]) > len(emotes_dictionary[j])
	})

	fmt.Print("\n--- Comenzando análisis ---\n")

	numero_comentarios_analizados := 0
	numero_lineas_leidas := 0
	startTime := time.Now()

	/* channel to query the api in several threads simultaneously */
	var wg sync.WaitGroup // Create a WaitGroup
	resultChan := make(chan WorkerResult)
	workers_added := 0
	toxicity_results := make([]LineResult, 0)
	current_index := 0

	// Read and process each line of the CSV file
	for {
		// Read one line from the CSV file
		record, err := reader.Read()

		for numero_lineas_leidas < num_existing_lines {
			numero_lineas_leidas++
			record, err = reader.Read()
		}

		// Check for end of file
		if err != nil {
			// If end of file, break out of the loop
			if err.Error() == "EOF" {
				break
			}
			// If any other error, log it and continue
			log.Printf("Error leyendo fila: %s", err)
			continue
		}

		// Process the record (each field is a slice element)
		//var toxicity_score float64
		for i, field := range record {
			if(i == 1){

				current_index++

				// Add default value to array of results
				// It will be overwriten if the comment is analyzed
				toxicity_results = append(toxicity_results, LineResult{record[0], record[1], "Omitido"})
		
				if !omit_comment(field, emotes_dictionary) {
					wg.Add(1)
					workers_added++
					go worker(field, current_index, &wg, resultChan)
				} 

				if workers_added >= num_workers {

					workers_added = 0
					numero_comentarios_analizados += workers_added
					//fmt.Println("***waiting***")

					// Wait for all goroutines to finish
					go func() {
						wg.Wait()
						close(resultChan) // Close the result channel after all goroutines are done
					}()
					// Receive and process results
					for result := range resultChan {
						toxicity_string := strconv.FormatFloat(result.Score * 100, 'f', 2, 64)
						toxicity_results[result.Index-1].Score = toxicity_string
						numero_comentarios_analizados++
					}

					//fmt.Println("***result***")
					//fmt.Println(toxicity_results)

					for _, lineResult := range toxicity_results {
	
						err := writer.Write([]string{lineResult.VideoId, lineResult.Comment, lineResult.Score})
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
					}

					if(numero_comentarios_analizados > 0 && numero_comentarios_analizados%100 == 0){
						executionTime := time.Since(startTime)
						fmt.Printf("%d Comentarios analizados. Tiempo: %s segundos\n", numero_comentarios_analizados, executionTime)
					}

					resultChan = make(chan WorkerResult)
					toxicity_results = make([]LineResult, 0)
					current_index = 0
					writer.Flush()
				}
				
				numero_lineas_leidas++
			}
		}

		if(numero_lineas_leidas > 0 && numero_lineas_leidas%1000 == 0){
			executionTime := time.Since(startTime)
			fmt.Printf("--- Línea %d. Tiempo: %s segundos ---\n", numero_lineas_leidas, executionTime)
		}

	}


}

func worker(comentario string, index int, wg *sync.WaitGroup, resultChan chan<- WorkerResult) {
	defer wg.Done() // Decrease the counter when the goroutine completes
	for {
		//fmt.Printf("analizando: %s\n", comentario)
		toxicity_score, err := analyze_comment(comentario)
		if err == nil {
			//fmt.Printf("fin %s. Toxicidad: %f\n", comentario, toxicity_score)
			resultChan <- WorkerResult{index, toxicity_score}
			return
		}

		fmt.Println(err)
		if err == ErrQuotaExceeded {
			time.Sleep(time.Second)
		}		
	}
	
}

func analyze_comment(comentario string) (float64, error){
	
	// Your Perspective API key
	envFile, _ := godotenv.Read(".env")
	apiKey := envFile["API_KEY"]
	if apiKey == "" {
		log.Fatal("Configura la API KEY en el fichero .env")
	}

	// API endpoint
	apiEndpoint := "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=" + apiKey

	// Data to be sent in the request
	data := map[string]interface{}{
		"comment": map[string]string{
			"text": comentario,
		},
		"languages": []string{
			"es", "en",
		},
		"requestedAttributes": map[string]map[string]interface{}{
			"TOXICITY": {},
		},
	}

	// Convert data to JSON format
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return 0, fmt.Errorf("error marshalling JSON:")
	}

	// Send HTTP request
	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		return 0, fmt.Errorf("error sending request: ", err)
		
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: ", err)
	}

	if strings.Contains(string(body), "RATE_LIMIT_EXCEEDED") {
		return 0, ErrQuotaExceeded
	}

	// Decode response JSON
	var response Score
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, fmt.Errorf("error decoding JSON: ", err)
	}

	// Get toxicity score
	//toxicityScore := response["attributeScores"].(map[string]interface{})["TOXICITY"].(map[string]interface{})["summaryScore"].(map[string]interface{})["value"].(float64)
	return response.AttributeScores.Toxicity.SummaryScore.Value, nil
}

func omit_comment(comment string, dictionary []string) bool{

	for _, val := range dictionary {
		// Replace the value with an empty string. 
		comment = strings.Replace(strings.ToLower(comment), strings.ToLower(val), "", -1)
	}

	var result []rune
	// Iterate over each character in the string
	for _, char := range comment {
		// Check if the character is printable
		if unicode.IsPrint(char) {
			// If it's printable, add it to the result slice
			result = append(result, char)
		}
	}
	comment = string(result)

	comment = strings.Replace(comment, " ", "", -1)
	if len(comment) == 0{
		return true
	}

	/* check if comment is composed only by one single character repeated */
	// Get the first character of the string
    firstChar := rune(comment[0])

    // Compare each character of the string with the first character
	single_char := true
    for _, char := range comment {
        if char != firstChar {
			single_char = false
            break
        }
    }

	return single_char

}

func get_num_existing_lines(output_csv string) int {
	// Open the CSV file for reading
	file, err := os.Open(output_csv)
	if err != nil {
		return -1
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Initialize line count
	lineCount := 0

	// Read the file line by line
	for {
		_, err := reader.Read()
		if err != nil {
			// Check if we reached end of file
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error:", err)
			return 0
		}
		// Increment line count
		lineCount++
	}

	return lineCount
}

type Score struct {
	AttributeScores AttributeScores `json:"attributeScores"`
}
type AttributeScores struct {
	Toxicity Toxicity `json:"TOXICITY"`
}
type Toxicity struct {
	SummaryScore SummaryScore `json:"summaryScore"`
}
type SummaryScore struct {
	Value float64 `json:"value"`
	Type string `json:"type"`
}
type WorkerResult struct {
	Index int
	Score float64
}
type LineResult struct {
	VideoId string
	Comment string
	Score string
}
var (
	ErrQuotaExceeded = errors.New("cuota excedida. Esperando")
)