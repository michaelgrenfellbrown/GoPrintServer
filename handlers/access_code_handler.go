package handlers

import (
	"html/template"
	"net/http"
	"printserver/config"
)

// AccessCodeHandler serves the page to display and print the daily access code
func AccessCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Load the access code from the configuration
	accessCode := config.AppConfig.AccessCode

	// Parse and render the template
	tmpl, err := template.ParseFiles("templates/access_code.html")
	if err != nil {
		http.Error(w, "Error loading access code template", http.StatusInternalServerError)
		return
	}

	// Pass the access code to the template
	data := map[string]interface{}{
		"AccessCode": accessCode,
	}
	tmpl.Execute(w, data)
}
