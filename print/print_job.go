package print

import (
	"log"
	"os/exec"
)

func PrintDocument(filePath, guestName string, tableNumber int) {
	log.Printf("Printing document for %s at table %d: %s", guestName, tableNumber, filePath)

	// Simulating print by using a command like "lp"
	cmd := exec.Command("lp", filePath)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to print document: %v", err)
	}
}
