package main

import (
	"assignment-3/controllers"
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Aplikasi berjalan pada http://localhost:8080")
	
	// Set interval to 15 seconds
	interval := 15 * time.Second

	// Start ticker for updating JSON file
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Start HTTP server
	http.HandleFunc("/", controllers.StatusHandler)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	// Update JSON file at regular intervals
	for range ticker.C {
		updateJSON()
	}
}

func updateJSON() {
	// Generate random values for water and wind
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	// Create status struct
	status := models.Status{
		Water: water,
		Wind:  wind,
	}

	// Marshal status struct to JSON
	jsonData, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		fmt.Printf("JSON marshaling error: %v\n", err)
		return
	}

	// Write JSON data to file
	err = os.WriteFile("status.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("File write error: %v\n", err)
		return
	}
}