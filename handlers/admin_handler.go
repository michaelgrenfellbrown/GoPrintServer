package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"printserver/config"
	"strconv"
)

// AdminHandler serves the admin page
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/admin.html")
		if err != nil {
			http.Error(w, "Error loading admin template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, config.AppConfig)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		config.AppConfig.ShopName = r.FormValue("shopName")
		config.AppConfig.ShopAddress = r.FormValue("shopAddress")
		config.AppConfig.PrinterURI = r.FormValue("printerURI")
		config.AppConfig.LogoPath = r.FormValue("logoPath")

		costPerPage, err := strconv.ParseFloat(r.FormValue("costPerPage"), 64)
		if err != nil {
			http.Error(w, "Invalid cost per page", http.StatusBadRequest)
			return
		}
		config.AppConfig.CostPerPage = costPerPage

		if err := config.SaveConfig(); err != nil {
			http.Error(w, "Error saving configuration", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

/*
*
// parseFloat is a helper function to parse form values into float64

	func parseFloat(value string, defaultValue float64) float64 {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return defaultValue
		}
		return parsed
	}

*
*/
func RegenerateAccessCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	newCode := config.RegenerateAccessCode()
	w.Write([]byte("New Access Code: " + newCode))

	// Update the access code in AppConfig
	config.AppConfig.AccessCode = newCode

	// Save the updated AppConfig to config.json
	err := config.SaveConfig()
	if err != nil {
		http.Error(w, "Failed to save access code", http.StatusInternalServerError)
		return
	}

	// Respond with the new access code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"newAccessCode": newCode,
	})
}
