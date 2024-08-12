package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

// type Location struct {
// 	Location []string `json:"locations"`
// }

var (
	Artists      []models.Artist
	Locations    []models.Location
	Dates        []models.Date
	Relations    []models.Relation
	locationMap  map[string]json.RawMessage
	dateMap      map[string]json.RawMessage
	realationMap map[string]json.RawMessage
)

func CollectData() []models.Data {
	FetchLocation()
	FetchArtists()
	FetchDate()
	FetchRelationData()

	data := make([]models.Data, len(Artists))

	for i := range Artists {
		data[i].A = Artists[i]
		data[i].D = Dates[i]
		data[i].L = Locations[i]
		data[i].R = Relations[i]
	}
	return data
}

func FetchLocation() []models.Location {
	location, err1 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer location.Body.Close()

	// locationData, err2 := io.ReadAll(location.Body)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	json.NewDecoder(location.Body).Decode(&dateMap)

	// println(string(locationData))
	// err3 := json.Unmarshal(locationData, &locationMap)
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }

	var bytes []byte
	for _, b := range locationMap {
		bytes = append(bytes, b...)
	}

	err4 := json.Unmarshal(bytes, &Locations)
	if err4 != nil {
		log.Fatal(err4)
	}
	return Locations
}

func FetchDate() []models.Date {
	date, err1 := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer date.Body.Close()

	// datesData, err2 := ioutil.ReadAll(date.Body)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	json.NewDecoder(date.Body).Decode(&dateMap)

	// err3 := json.Unmarshal(datesData, &dateMap)
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }

	var bytes []byte
	for _, b := range dateMap {
		bytes = append(bytes, b...)
	}

	err4 := json.Unmarshal(bytes, &Dates)
	if err4 != nil {
		log.Fatal(err4)
	}
	return Dates
}

func FetchRelationData() []models.Relation {
	relation, err1 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err1 != nil {
		log.Fatal(err1)
	}

	defer relation.Body.Close()

	// relationData, err2 := io.ReadAll(relation.Body)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	// err3 := json.Unmarshal(relationData, &realationMap)
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }
	json.NewDecoder(relation.Body).Decode(&dateMap)

	var bytes []byte
	for _, b := range realationMap {
		bytes = append(bytes, b...)
	}

	err4 := json.Unmarshal(bytes, &Relations)
	if err4 != nil {
		log.Fatal(err4)
	}
	return Relations
}

func FetchArtists() []models.Artist {
	resArtist, err1 := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resArtist.Body.Close()

	artsisInfo, err2 := io.ReadAll(resArtist.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := json.Unmarshal(artsisInfo, &Artists)
	if err3 != nil {
		log.Fatal(err3)
	}
	return Artists
}
