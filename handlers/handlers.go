package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
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

func StaticServer(w http.ResponseWriter, r *http.Request) {
	filePath := "." + r.URL.Path
	info, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if info.IsDir() {
		// The path is a directory, return a 404
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/artistInfo" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("./templates/artistPage.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	artistName := r.FormValue("ArtistName")

	if artistName == "" {
		http.Error(w, "400 Bad Request: Missing artist name", http.StatusBadRequest)
		return
	}

	artistInfo := api.CollectData() // Assuming this returns []models.Data

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
		http.Error(w, "404 Artist Not Found", http.StatusNotFound)
		return
	}

	if err := t.Execute(w, art); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}
