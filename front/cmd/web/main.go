package main

import (
	"ERGSurvey/back/app/survey"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//port := os.Getenv("ROOT_URL")
	port := "8080"

	// Make the GET request
	resp, err := http.Get("http://back-service:8081/latest")
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderQuestion(w, "templates/survey.component.gohtml", &question)
	})

	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		renderQuestion(w, "templates/table.component.gohtml", &question)
	})

	fmt.Printf("Starting survey frontend on port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}
