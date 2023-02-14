package main

import (
	"ERGSurvey/back/app/survey"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getLatestSurveyData() *survey.Question {
	// Make the GET request
	resp, err := http.Get("https://back-service:8081/latest")
	if err != nil {
		fmt.Printf("failed to make the GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	reader := bufio.NewReader(resp.Body)
	decoder := json.NewDecoder(reader)
	var question survey.Question
	err = decoder.Decode(&question)
	if err != nil {
		fmt.Printf("failed to unmarshal the response body: %v", err)
	}

	return &question
}

func main() {
	port := "8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderQuestion(w, "templates/survey.component.gohtml", getLatestSurveyData())
	})

	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		renderQuestion(w, "templates/table.component.gohtml", getLatestSurveyData())
	})

	fmt.Printf("Starting survey frontend on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}
