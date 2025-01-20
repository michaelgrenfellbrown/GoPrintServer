package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"printserver/config"
	"printserver/print"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the upload page (HTML form)
		http.ServeFile(w, r, "templates/upload.html")
	case http.MethodPost:
		// Handle file upload
		// (Existing file upload logic here)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	guestName := r.FormValue("guestName")
	tableNumber, err := strconv.Atoi(r.FormValue("tableNumber"))
	if err != nil {
		http.Error(w, "Invalid table number", http.StatusBadRequest)
		return
	}

	// Access code validation
	accessCode := r.FormValue("accessCode")
	if accessCode != config.GetAccessCode() {
		http.Error(w, "Invalid access code", http.StatusForbidden)
		return
	}

	// File upload handling
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading uploaded file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Prepare the print job
	numPages := 5 // Placeholder, replace with actual page count calculation
	totalCost := float64(numPages) * config.GetConfig().CostPerPage
	printJob := print.PrintJob{
		GuestName:   guestName,
		TableNumber: tableNumber,
		FileName:    handler.Filename,
		NumPages:    numPages,
		TotalCost:   totalCost,
		DateTime:    time.Now(),
	}

	// Process the print job
	err = print.ProcessPrintJob(printJob, filePath)
	if err != nil {
		http.Error(w, "Error processing print job", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Print job submitted successfully"))
}
