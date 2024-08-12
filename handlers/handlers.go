package handlers

import (
	"encoding/json"
	"strings"

	"html/template"
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(api.FetchArtists())
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(api.FetchLocation())
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(api.FetchDate())
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(api.FetchRelationData()); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	if strings.ToUpper(r.Method) != "GET" {
		http.Error(w, "400 Page not found", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./templates/Home.html")
	if err != nil {
		log.Fatal(err)
	}

	data := api.FetchArtists()

	tmp.Execute(w, data)

}
