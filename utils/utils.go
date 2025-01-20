package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

// CalculatePages determines the number of pages in a PDF file
// Requires a utility like `pdfinfo` from poppler-utils
func CalculatePages(filePath string) int {
	// Command to get PDF information
	cmd := exec.Command("pdfinfo", filePath)
	output, err := cmd.Output()
	if err != nil {
		// Default to 1 page if the page count can't be determined
		return 1
	}

	// Parse the output to find the page count
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Pages:") {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				if pages, err := strconv.Atoi(parts[1]); err == nil {
					return pages
				}
			}
		}
	}

	// Default to 1 page if parsing fails
	return 1
}
