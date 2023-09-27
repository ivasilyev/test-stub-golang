package main

import (
	"fmt"
	"log"
	"net/http"
	"src/main/go/org/test/golang_stub/main.go/src/main/go/org/test/golang_stub/constants"
	"src/main/go/org/test/golang_stub/main.go/src/main/go/org/test/golang_stub/controller"
)

const PORT = 10100

func main() {
	// Define the endpoint HTTP routes
	controller.HandleRestEndpoint()
	controller.HandleWebEndpoint()

	// Define the common HTTP routes
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(constants.CSS_DIR))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(constants.JS_DIR))))

	// Start the HTTP server
	log.Println(fmt.Sprintf("Server started on http://localhost:%d", PORT))
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
