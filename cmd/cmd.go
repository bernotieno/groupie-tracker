package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/handlers"
	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

func Exec() {
	if len(os.Args) != 1 {
		fmt.Println("USAGE: go run main.go")
		return
	}
	// Collect and preload data
	data, err := api.CollectData()
	if err != nil {
		log.Fatalf("Failed to collect data: %v", err)
	}
	// Initialize the search index
	searchIndex := &models.SearchIndex{
		ArtistName:   make(map[string][]models.IndexedData),
		MemberName:   make(map[string][]models.IndexedData),
		LocationName: make(map[string][]models.IndexedData),
		FirstAlbum:   make(map[string][]models.IndexedData),
		CreationDate: make(map[int][]models.IndexedData),
	}

	searchIndex.PreloadData(data)

	http.HandleFunc("/artistInfo", handlers.ArtistInfo)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/static/", handlers.StaticServer)
	http.HandleFunc("/search", searchIndex.SearchHandler)
	fmt.Println("http://localhost:8081/")
	http.ListenAndServe(":8081", nil)
}
