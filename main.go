package main

import (
	"fmt"
	"net/http"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/artists", handlers.ArtistsHandler)
	http.HandleFunc("/locations", handlers.LocationsHandler)
	http.HandleFunc("/dates", handlers.DatesHandler)
	http.HandleFunc("/relations", handlers.RelationsHandler)
	http.HandleFunc("/artistInfo", handlers.ArtistInfo)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/static/", handlers.StaticServer)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
