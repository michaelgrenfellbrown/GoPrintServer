package config

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

var (
	accessCode string
	mu         sync.Mutex
)

// Initialize the access code during package initialization
func init() {
	// rand.Seed(time.Now().UnixNano())
	accessCode = generateAccessCode()
}

// generateAccessCode creates a new 5-digit random access code
func generateAccessCode() string {
	return fmt.Sprintf("%05d", rand.Intn(100000))
}

// GetAccessCode returns the current access code
func GetAccessCode() string {
	mu.Lock()
	defer mu.Unlock()
	return accessCode
}

// RegenerateAccessCode generates and returns a new access code
func RegenerateAccessCode() string {
	mu.Lock()
	defer mu.Unlock()
	accessCode = generateAccessCode()
	return accessCode
}

// Config structure
type Config struct {
	ShopName    string  `json:"ShopName"`
	ShopAddress string  `json:"ShopAddress"`
	CostPerPage float64 `json:"CostPerPage"`
	PrinterURI  string  `json:"PrinterURI"`
	AccessCode  string  `json:"AccessCode"`
}

// Global AppConfig instance
var AppConfig Config

// LoadConfig loads the application configuration from a JSON file
func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}
	return nil
}

// SaveConfig saves the current configuration to the JSON file
func SaveConfig() error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print the JSON
	return encoder.Encode(AppConfig)
}

// GetConfig returns the loaded application configuration
func GetConfig() Config {
	return AppConfig
}
