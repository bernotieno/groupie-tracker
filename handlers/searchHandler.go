package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

// NewSearchIndex builds and returns a new SearchIndex
func NewSearchIndex(data []models.Data) *models.SearchIndex {
	index := &models.SearchIndex{
		ArtistName:   make(map[string][]models.IndexedData),
		MemberName:   make(map[string][]models.IndexedData),
		LocationName: make(map[string][]models.IndexedData),
		FirstAlbum:   make(map[string][]models.IndexedData),
		CreationDate: make(map[int][]models.IndexedData),
	}

	for _, d := range data {
		artistName := strings.ToLower(d.A.Name)
		indexedData := models.IndexedData{
			Data:       d,
			ArtistName: d.A.Name,
		}

		// Index by artist name
		index.ArtistName[artistName] = append(index.ArtistName[artistName], indexedData)

		// Index by member names
		for _, member := range d.A.Members {
			index.MemberName[strings.ToLower(member)] = append(index.MemberName[strings.ToLower(member)], indexedData)
		}

		// Index by location names
		for _, location := range d.L.Locations {
			index.LocationName[strings.ToLower(location)] = append(index.LocationName[strings.ToLower(location)], indexedData)
		}

		// Index by first album name
		index.FirstAlbum[strings.ToLower(d.A.FirstAlbum)] = append(index.FirstAlbum[strings.ToLower(d.A.FirstAlbum)], indexedData)

		// Index by creation date
		index.CreationDate[d.A.CreationDate] = append(index.CreationDate[d.A.CreationDate], indexedData)
	}

	fmt.Println("New search index created.")
	return index
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is correct
	if r.URL.Path != "/search" {
		HandleError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	// Check if the request method is correct (if needed)
	// if strings.ToUpper(r.Method) != "POST" {
	//     HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	//     return
	// }

	fmt.Println("Called")

	query := r.URL.Query().Get("q")

	// Check internet connection
	msg := checkInternetConnection()
	if msg != "" {
		fmt.Println(msg)
		fmt.Println("HERE ERROR PAGE IS SUPPOSED TO BE DISPLAYED")
		HandleError(w, "Internet Connectivity Issues", http.StatusRequestTimeout)
		return
	}

	// Collect data from the API
	artistsData, err := api.CollectData()
	if err != nil {
		HandleError(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a search index and perform the search
	index := NewSearchIndex(artistsData)
	results := index.Search(query)
	// fmt.Println(results)
	// Send results as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		HandleError(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
