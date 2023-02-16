package main

import (
	"ERGSurvey/board/anonimousboard"
	"fmt"
	"log"
	"net/http"
)

type BoardQuestionPayload struct {
	Question string
}

func main() {
	port := "8082"

	board := anonimousboard.AnonBoard{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderQuestions(w, "templates/board.component.gohtml", board.GetAllQuestions())
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Printf("Method submit has been called\n")
			res := &BoardQuestionPayload{}
			err := readJSON(w, r, res)
			if err != nil {
				fmt.Printf("Falsed to parse JSON: %v\n", err)
			}
			ip := readUserIP(r)
			fmt.Printf("Request from: %s\n", ip)
			board.AddQuestion(res.Question)
		}
	})

	fmt.Printf("Starting survey frontend on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}
