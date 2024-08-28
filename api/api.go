package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

var (
	locationMap  map[string]json.RawMessage
	dateMap      map[string]json.RawMessage
	realationMap map[string]json.RawMessage
)

// CollectData fetches and compiles artist, location, date, and relation data 
// from predefined API endpoints into a slice of `models.Data` objects. It logs 
// and returns errors if any of the fetch operations fail.
func CollectData() ([]models.Data, error) {
	Locations, err1 := FetchLocation("https://groupietrackers.herokuapp.com/api/locations")
	if err1 != nil {
		log.Println(err1)
		return []models.Data{}, err1
	}
	Artists, err2 := FetchArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err2 != nil {
		log.Println(err2)
		return []models.Data{}, err2
	}
	Dates, err3 := FetchDate("https://groupietrackers.herokuapp.com/api/dates")
	if err3 != nil {
		log.Println(err3)
		return []models.Data{}, err3
	}
	Relations, err4 := FetchRelationData("https://groupietrackers.herokuapp.com/api/relation")
	if err4 != nil {
		log.Println(err4)
		return []models.Data{}, err4
	}

	data := make([]models.Data, len(Artists))

	for i := range Artists {
		data[i].A = Artists[i]
		data[i].D = Dates[i]
		data[i].L = Locations[i]
		data[i].R = Relations[i]
	}
	return data, nil
}

// FetchLocation retrieves location data from the specified URL, processes the JSON response,
// and returns a slice of location models. It logs and returns errors if the HTTP request,
// reading the response body, or JSON unmarshalling fails.
func FetchLocation(url string) ([]models.Location, error) {
	location, err1 := http.Get(url)
	if err1 != nil {
		log.Println(err1)
		return []models.Location{}, err1
	}
	defer location.Body.Close()

	locationData, err2 := io.ReadAll(location.Body)
	if err2 != nil {
		log.Println(err2)
		return []models.Location{}, err2
	}

	err3 := json.Unmarshal(locationData, &locationMap)
	if err3 != nil {
		log.Println(err3)
		return []models.Location{}, err3
	}

	var bytes []byte
	for _, b := range locationMap {
		bytes = append(bytes, b...)
	}
	var Locations []models.Location
	err4 := json.Unmarshal(bytes, &Locations)
	if err4 != nil {
		log.Println(err4)
		return []models.Location{}, err4
	}
	return Locations, nil
}

// FetchDate retrieves date data from the specified URL, processes the JSON response,
// and returns a slice of date models. It logs and returns errors if the HTTP request,
// reading the response body, or JSON unmarshalling fails.
func FetchDate(url string) ([]models.Date, error) {
	date, err1 := http.Get(url)
	if err1 != nil {
		log.Println(err1)
		return []models.Date{}, err1
	}
	defer date.Body.Close()

	datesData, err2 := io.ReadAll(date.Body)
	if err2 != nil {
		log.Println(err2)
		return []models.Date{}, err2
	}

	err3 := json.Unmarshal(datesData, &dateMap)
	if err3 != nil {
		log.Println(err3)
		return []models.Date{}, err3
	}

	var bytes []byte
	for _, b := range dateMap {
		bytes = append(bytes, b...)
	}
	var Dates []models.Date
	err4 := json.Unmarshal(bytes, &Dates)
	if err4 != nil {
		log.Println(err4)
		return []models.Date{}, err4
	}
	return Dates, nil
}

// FetchRelationData retrieves relation data from the specified URL, processes the JSON response,
// and returns a slice of relation models. It logs and returns errors if the HTTP request,
// reading the response body, or JSON unmarshalling fails.
func FetchRelationData(url string) ([]models.Relation, error) {
	relation, err1 := http.Get(url)
	if err1 != nil {
		log.Println(err1)
		return []models.Relation{}, err1
	}

	defer relation.Body.Close()

	relationData, err2 := io.ReadAll(relation.Body)
	if err2 != nil {
		log.Println(err2)
		return []models.Relation{}, err2
	}

	err3 := json.Unmarshal(relationData, &realationMap)
	if err3 != nil {
		log.Println(err3)
		return []models.Relation{}, err3
	}

	var bytes []byte
	for _, b := range realationMap {
		bytes = append(bytes, b...)
	}
	var Relations []models.Relation
	err4 := json.Unmarshal(bytes, &Relations)
	if err4 != nil {
		log.Println(err4)
		return []models.Relation{}, err4
	}
	return Relations, nil
}

// FetchArtists retrieves artist data from the specified URL, parses the JSON response,
// and returns a slice of artist models. It logs and returns errors if the HTTP request,
// reading response body, or JSON unmarshalling fails.
func FetchArtists(url string) ([]models.Artist, error) {
	resArtist, err1 := http.Get(url)
	if err1 != nil {
		log.Println(err1)
		return []models.Artist{}, err1
	}
	defer resArtist.Body.Close()

	artsisInfo, err2 := io.ReadAll(resArtist.Body)
	if err2 != nil {
		log.Println(err2)
		return []models.Artist{}, err2
	}
	var Artists []models.Artist
	err3 := json.Unmarshal(artsisInfo, &Artists)
	if err3 != nil {
		log.Println(err3)
		return []models.Artist{}, err3
	}
	return Artists, nil
}
