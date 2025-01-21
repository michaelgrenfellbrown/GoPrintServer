package main

import (
	"log"
	"net/http"
	"printserver/config"
	"printserver/handlers"
	"printserver/print"
)

func main() {
	// Load configuration
	err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Static file server for styles and images
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers
	http.HandleFunc("/", handlers.UploadHandler)
	http.HandleFunc("/admin", handlers.AdminHandler)
	http.HandleFunc("/admin/regenerate-code", handlers.RegenerateAccessCodeHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/print-access-code", handlers.AccessCodeHandler)

	// Initialize the printer service
	print.InitializePrinterService()

	// Start the server
	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
