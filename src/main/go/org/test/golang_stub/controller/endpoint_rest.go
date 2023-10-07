package controller

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"src/main/go/org/test/golang_stub/main.go/src/main/go/org/test/golang_stub/constants"
	"src/main/go/org/test/golang_stub/main.go/src/main/go/org/test/golang_stub/dto"
)

const ENDPOINT_NAME = "endpoint"
const DEFAULT_DELAY_MS = 500

type TemplateData struct {
	EndpointName string
}

// Define the default Delay
var delayDto = dto.DelayDto{
	DelayMs: DEFAULT_DELAY_MS,
}

// Counter to be monitored
var opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_processed_ops_total",
	Help: "The total number of processed events",
})

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
	var newDelayDto dto.DelayDto

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

/*
   curl \
       --request GET \
       "http://localhost:10000/endpoint"
*/
func endPointHandler(responseWriter http.ResponseWriter, requests *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	// Response code here
	responseWriter.Write([]byte("{\"success\":true}"))
	opsProcessed.Inc()
}

func HandleRestEndpoint() {
	http.HandleFunc(fmt.Sprintf("/%s", ENDPOINT_NAME), endPointHandler)
	http.HandleFunc(fmt.Sprintf("/%s%s", ENDPOINT_NAME, constants.API_GET_SUFFIX), getDelayHandler)
	http.HandleFunc(fmt.Sprintf("/%s%s", ENDPOINT_NAME, constants.API_SET_SUFFIX), setDelayHandler)
	http.HandleFunc(fmt.Sprintf("/%s%s", ENDPOINT_NAME, constants.API_RESET_SUFFIX), resetDelayHandler)
}
