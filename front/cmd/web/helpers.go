package main

import (
	"ERGSurvey/back/app/survey"
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates
var templateFS embed.FS

func renderQuestion(w http.ResponseWriter, t string, question *survey.Question) {
	tmpl, err := template.ParseFS(templateFS, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, *question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
