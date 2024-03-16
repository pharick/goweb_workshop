package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, name string, data map[string]any) {
	// find all templates in the templates directory
	templates, err := filepath.Glob("templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// read all templates from the disk
	t, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send the template to the client
	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}
