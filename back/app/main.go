package main

import (
	"ERGSurvey/back/app/survey"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8081"

	surv := survey.CreateDummySurvey()

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Printf("Method submit has been called\n")
			res := &survey.ResponsePayload{}
			err := readJSON(w, r, res)
			if err != nil {
				fmt.Printf("Falsed to parse JSON: %v\n", err)
			}
			surv.Increment(res.Id)
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

	fmt.Printf("Starting survey backen on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}
