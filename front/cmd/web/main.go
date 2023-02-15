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

	return &question
}

func main() {
	port := "8080"

	surv := survey.CreateDummySurvey()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderQuestion(w, "templates/survey.component.gohtml", surv.LatestQuestion())
	})

	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		renderQuestion(w, "templates/table.component.gohtml", surv.LatestQuestion())
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Printf("Method submit has been called\n")
			res := &survey.ResponsePayload{}
			err := readJSON(w, r, res)
			if err != nil {
				fmt.Printf("Falsed to parse JSON: %v\n", err)
			}
			ip := readUserIP(r)
			fmt.Printf("Request from: %s\n", ip)
			surv.Increment(res.Id, ip)
		}
	})

	http.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Printf("Method latest has been called\n")
			data, err := json.Marshal(*surv.LatestQuestion())
			if err != nil {
				fmt.Printf("failed to encode the object to JSON: %v", err)
			}
			w.Write(data)
		}
	})

	fmt.Printf("Starting survey frontend on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}
