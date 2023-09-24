package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type MyWebpage struct {
	IntValue int `json:"intValue"`
}

var webpage MyWebpage

func main() {
	// Define the default value
	webpage = MyWebpage{
		IntValue: 0,
	}

	// Define the HTTP routes
	http.HandleFunc("/api/value", getValueHandler)
	http.HandleFunc("/api/value/set", setValueHandler)
	http.HandleFunc("/api/value/reset", resetValueHandler)

	// Start the HTTP server
	log.Println("Server started on http://localhost:10100")
	log.Fatal(http.ListenAndServe(":10100", nil))
}

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the MyWebpage struct to JSON
	jsonValue, _ := json.Marshal(webpage)

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonValue)
}

func setValueHandler(w http.ResponseWriter, r *http.Request) {
	newValueStr := r.FormValue("value")

	newValue, err := strconv.Atoi(newValueStr)
	if err != nil {
		http.Error(w, "Invalid integer value", http.StatusBadRequest)
		return
	}

	// Set the new integer value
	webpage.IntValue = newValue

	// Convert the MyWebpage struct to JSON
	jsonValue, _ := json.Marshal(webpage)

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonValue)
}

func resetValueHandler(w http.ResponseWriter, r *http.Request) {
	// Reset the integer value to the default
	webpage.IntValue = 0

	// Convert the MyWebpage struct to JSON
	jsonValue, _ := json.Marshal(webpage)

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonValue)
}
