package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Open the CSV file
	file, err := os.Open("data_frame_twitch_mensajes.csv")
	if err != nil {
		log.Fatalf("Error abriendo csv: %s. ¿Está en el mismo directorio que el ejecutable?", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	numero_comentarios_analizados := 0
	startTime := time.Now()

	// Read and process each line of the CSV file
	for {
		// Read one line from the CSV file
		record, err := reader.Read()

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
		for i, field := range record {
			if(i == 1){
				analizar_comentario(field)
				numero_comentarios_analizados++
			}
		}

		if(numero_comentarios_analizados%10 == 0){
			executionTime := time.Now().Sub(startTime)
			fmt.Printf("%d Comentarios analizados. Tiempo: %s segundos\n", numero_comentarios_analizados, executionTime)
		}

	}


}

func analizar_comentario(comentario string){
	
	// Your Perspective API key
	apiKey := "AIzaSyDSnvUYyHipzZC1A_Xl2rg325Ys0V1tYvU"

	// API endpoint
	apiEndpoint := "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=" + apiKey

	// Data to be sent in the request
	data := map[string]interface{}{
		"comment": map[string]string{
			"text": comentario,
		},
		"requestedAttributes": map[string]map[string]interface{}{
			"TOXICITY": {},
		},
	}

	// Convert data to JSON format
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send HTTP request
	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Decode response JSON
	var response Score
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//fmt.Println(response);

	// Get toxicity score
	//toxicityScore := response["attributeScores"].(map[string]interface{})["TOXICITY"].(map[string]interface{})["summaryScore"].(map[string]interface{})["value"].(float64)
	toxicityScore := response.AttributeScores.Toxicity.SummaryScore.Value

	// Output the toxicity score
	fmt.Printf("Comentario: %s. Toxicidad: %.2f%%\n", comentario, toxicityScore * 100)

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
	Value float32 `json:"value"`
	Type string `json:"type"`
}