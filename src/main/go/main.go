package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const PORT = 10100
const ENDPOINT_NAME = "endpoint"
const DEFAULT_DELAY_MS = 500

type TemplateData struct {
	EndpointName string
}

type DelayDto struct {
	DelayMs int `json:"delayMs"`
}

// Define the default Delay
var delayDto = DelayDto{
	DelayMs: DEFAULT_DELAY_MS,
}

/*
   curl \
       --request GET \
       "http://localhost:10100/endpoint/api/get-delay"
*/
func getDelayHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the MyWebpage struct to JSON
	jsonDelay, _ := json.Marshal(delayDto)

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDelay)
}

/*
   curl \
       --request POST \
       --data-binary '{"delayMs":555}' \
       --header 'Content-Type: application/json' \
       "http://localhost:10100/endpoint/api/set-delay"
*/
func setDelayHandler(responseWriter http.ResponseWriter, requests *http.Request) {
	var newDelayDto DelayDto

	decoder := json.NewDecoder(requests.Body)
	err := decoder.Decode(&newDelayDto)
	if err != nil {
		http.Error(responseWriter, "Invalid JSON", http.StatusBadRequest)
		return
	}
	delayDto.DelayMs = newDelayDto.DelayMs
	getDelayHandler(responseWriter, requests)
}

/*
   curl \
       --request POST \
       "http://localhost:10100/endpoint/api/reset-delay"
*/
func resetDelayHandler(w http.ResponseWriter, r *http.Request) {
	// Reset the integer Delay to the default
	delayDto.DelayMs = DEFAULT_DELAY_MS

	// Convert the MyWebpage struct to JSON
	jsonDelay, _ := json.Marshal(delayDto)

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDelay)
}

func endPointHandler(responseWriter http.ResponseWriter, requests *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write([]byte("{\"success\":true}"))
}

func renderTemplate(responseWriter http.ResponseWriter, templateData TemplateData) {
	tmpl, err := template.ParseFiles("../resources/static/templates/ui.html")
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(responseWriter, templateData)
}

func endPointUiHandler(responseWriter http.ResponseWriter, requests *http.Request) {
	renderTemplate(responseWriter, TemplateData{
		EndpointName: ENDPOINT_NAME,
	})
}

func main() {
	// Define the endpoint HTTP routes
	http.HandleFunc(fmt.Sprintf("/%s", ENDPOINT_NAME), endPointHandler)
	http.HandleFunc(fmt.Sprintf("/%s/ui", ENDPOINT_NAME), endPointUiHandler)
	http.HandleFunc(fmt.Sprintf("/%s/api/get-delay", ENDPOINT_NAME), getDelayHandler)
	http.HandleFunc(fmt.Sprintf("/%s/api/set-delay", ENDPOINT_NAME), setDelayHandler)
	http.HandleFunc(fmt.Sprintf("/%s/api/reset-delay", ENDPOINT_NAME), resetDelayHandler)

	// Define the common HTTP routes
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../resources/static/styles/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../resources/static/scripts/"))))

	// Start the HTTP server
	log.Println(fmt.Sprintf("Server started on http://localhost:%d", PORT))
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
