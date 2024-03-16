package controllers

import (
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	_ = r

	// Read status JSON file
	file, err := os.ReadFile("./status.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("File read error: %v", err), http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON data into status struct
	var status models.Status
	err = json.Unmarshal(file, &status)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON unmarshaling error: %v", err), http.StatusInternalServerError)
		return
	}

	// Determine status based on water and wind values
	var waterStatus, windStatus string
	if status.Water < 5 {
		waterStatus = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	if status.Wind < 6 {
		windStatus = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	// Prepare response
	water := fmt.Sprintf("Ketinggian Air: %d meter, Status : %s\n", status.Water, waterStatus)
	wind := fmt.Sprintf("Kecepatan Angin: %d meter per detik, Status : %s\n", status.Wind, windStatus)

	// Write response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
			<head>
				<meta http-equiv="refresh" content="15">
				<title>Laporan Cuaca</title>
				<style>
					body {
						display: flex;
						justify-content: center;
						align-items: center;
						text-align: center;
						height: 100vh;
						margin: 0;
					}
				</style>
			</head>
			<body>
				%s
				<br>
				%s
			</body>
		</html>`, water, wind)
}