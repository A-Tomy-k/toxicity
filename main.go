package main

import (
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
	"time"
	"unicode"

	"github.com/joho/godotenv"
)

const input_csv = "data_frame_twitch_mensajes.csv";
const output_csv = "resultado_toxicidad.csv";

func main() {
/*
	// Set up a signal handler to catch SIGINT (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	// Define a channel to indicate when the program should exit
	exitChan := make(chan bool)
*/
	num_existing_lines := get_num_existing_lines()
	fmt.Printf("CSV de salida con %d líneas ya analizadas\n", num_existing_lines)

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

	// Order dictionary to find longer emotes first. 
	// Example: rubbComo and rubbC collide, so we need to check first the longest
	sort.Slice(emotes_dictionary, func(i, j int) bool {
		return emotes_dictionary[i] > emotes_dictionary[j]
	})

	// Handle SIGINT signal asynchronously
	/*go func() {
		<-signalChan
		fmt.Println("\nReceived SIGINT. Cleaning up and exiting.")
		// Close the CSV file and flush the writer
		writer.Flush()
		outputFile.Close()
		exitChan <- true
	}()*/

	numero_comentarios_analizados := 0
	numero_lineas_leidas := 0
	startTime := time.Now()

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
		var toxicity_score float64
		for i, field := range record {
			if(i == 1){

				toxicity_score = -1
				if !omit_comment(field, emotes_dictionary) {
					for {
						toxicity_score, err = analyze_comment(field)
						if err == nil {
							break	
						}
	
						fmt.Println(err)
						if err == ErrQuotaExceeded {
							time.Sleep(time.Second)
						}
					}
				}

				toxicity_string := "Omitido"
				if toxicity_score != -1 {
					toxicity_string = strconv.FormatFloat(toxicity_score * 100, 'f', 2, 64)
				}

				err := writer.Write([]string{record[0], record[1], toxicity_string})
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				numero_comentarios_analizados++
			}
		}

		if(numero_comentarios_analizados > 0 && numero_comentarios_analizados%100 == 0){
			executionTime := time.Since(startTime)
			fmt.Printf("%d Comentarios analizados. Tiempo: %s segundos\n", numero_comentarios_analizados, executionTime)
		}

		if(numero_lineas_leidas > 0 && numero_lineas_leidas%1000 == 0){
			executionTime := time.Since(startTime)
			fmt.Printf("%d líneas leídas. Tiempo: %s segundos\n", numero_lineas_leidas, executionTime)
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

func get_num_existing_lines() int {
	// Open the CSV file for reading
	file, err := os.Open(output_csv)
	if err != nil {
		fmt.Println("Error contando líneas:", err)
		return 0
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
var (
	ErrQuotaExceeded = errors.New("cuota excedida. Esperando")
)