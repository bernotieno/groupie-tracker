package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
)

var artistUrl = "https://groupietrackers.herokuapp.com/api/artists"

// Home handles requests to the root path ("/") and serves the Home.html template.
// It checks for valid request methods and internet connectivity before fetching
// artist data and rendering the template. Errors are handled with appropriate
// HTTP status codes and messages.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if strings.ToUpper(r.Method) != "GET" {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./templates/Home.html")
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	if err := CheckInternetConnection(); err != nil {
		log.Println(err)
		HandleError(w, "Poor Internet Connectivity", http.StatusGatewayTimeout)
		return
	}

	data, err := api.FetchArtists(artistUrl)
	if err != nil {
		log.Println(err)
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err1 := tmp.Execute(w, data)
	if err1 != nil {
		log.Println(err1)
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
