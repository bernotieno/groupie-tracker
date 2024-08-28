package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)


// ArtistInfo handles POST requests to "/artistInfo" by rendering artist data 
// from a template. It returns errors for invalid methods, missing artist names, 
// poor connectivity, or if the artist is not found.
func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		HandleError(w, "400 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/artistInfo" {
		HandleError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("./templates/artistPage.html")
	if err != nil {
		HandleError(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	artistName := r.FormValue("ArtistName")

	if artistName == "" {
		HandleError(w, "Bad Request: Missing artist name", http.StatusBadRequest)
		return
	}
	msg := checkInternetConnection()
	if msg != "" {
		fmt.Println(msg)
		HandleError(w, "Poor Internet Connectivity", http.StatusRequestTimeout)
		return
	}

	artistInfo, err := api.CollectData()
	if err != nil {
		HandleError(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	var art models.Data
	found := false
	for _, artist := range artistInfo {
		if artistName == artist.A.Name {
			art = artist
			found = true
			break
		}
	}

	if !found {
		HandleError(w, "404 Artist Not Found", http.StatusNotFound)
		return
	}

	if err := t.Execute(w, art); err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}
