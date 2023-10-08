package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"src/main/go/org/test/golang_stub/main.go/src/main/go/org/test/golang_stub/constants"
)

type TemplateData struct {
	EndpointName string
}

func renderTemplate(responseWriter http.ResponseWriter, templateData TemplateData) {
	tmpl, err := template.ParseFiles(constants.UI_TEMPLATE_FILE)
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

func HandleWebEndpoint() {
	http.HandleFunc(fmt.Sprintf("/%s%s", ENDPOINT_NAME, constants.UI_SUFFIX), endPointUiHandler)
}
