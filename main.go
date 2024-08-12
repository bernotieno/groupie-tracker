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
	http.HandleFunc("/", handlers.Home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
