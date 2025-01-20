package handlers

import (
	"html/template"
	"net/http"
	"printserver/config"
	"strconv"
)

// AdminHandler serves the admin page
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the admin page with the current configuration
		tmpl, err := template.ParseFiles("templates/admin.html")
		if err != nil {
			http.Error(w, "Error loading admin template", http.StatusInternalServerError)
			return
		}

		// Pass the AppConfig fields to the template
		data := map[string]interface{}{
			"ShopName":    config.AppConfig.ShopName,
			"ShopAddress": config.AppConfig.ShopAddress,
			"CostPerPage": config.AppConfig.CostPerPage,
			"PrinterURI":  config.AppConfig.PrinterURI,
			"AccessCode":  config.AppConfig.AccessCode,
		}
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		// Handle form submissions to update config.json
		r.ParseForm()
		config.AppConfig.ShopName = r.FormValue("shopName")
		config.AppConfig.ShopAddress = r.FormValue("shopAddress")
		config.AppConfig.CostPerPage = parseFloat(r.FormValue("costPerPage"), 0.50)
		config.AppConfig.PrinterURI = r.FormValue("printerURI")
		config.AppConfig.AccessCode = r.FormValue("accessCode")

		err := config.SaveConfig()
		if err != nil {
			http.Error(w, "Failed to save config", http.StatusInternalServerError)
			return
		}

		// Redirect back to the admin page
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

// parseFloat is a helper function to parse form values into float64
func parseFloat(value string, defaultValue float64) float64 {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return parsed
}

func RegenerateAccessCodeHandler(w http.ResponseWriter, r *http.Request) {
	newCode := config.RegenerateAccessCode()
	w.Write([]byte("New Access Code: " + newCode))
}
