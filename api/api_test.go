package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

func TestFetchArtists(t *testing.T) {
	mockArtistsData := []models.Artist{}
	// Set up a mock server to handle the API request
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare the mock response data
		mockArtistsData = []models.Artist{
			{
				ID:           1,
				Name:         "Artist 1",
				Image:        "https://example.com/artist1.jpg",
				CreationDate: 2010,
				FirstAlbum:   "First Album",
				Members:      []string{"Member 1", "Member 2"},
			},
			{
				ID:           2,
				Name:         "Artist 2",
				Image:        "https://example.com/artist2.jpg",
				CreationDate: 2015,
				FirstAlbum:   "Second Album",
				Members:      []string{"Member 3", "Member 4"},
			},
			{
				ID:           3,
				Name:         "Artist 3",
				Image:        "https://example.com/artist3.jpg",
				CreationDate: 2020,
				FirstAlbum:   "Third Album",
				Members:      []string{"Member 5", "Member 6"},
			},
		}

		// Encode the mock data as JSON
		mockArtistsDataBytes, err := json.Marshal(mockArtistsData)
		if err != nil {
			t.Errorf("Failed to marshal mock artists data: %v", err)
			return
		}

		// Write the mock response
		w.Write(mockArtistsDataBytes)
	}))
	defer mockServer.Close()

	// Call the FetchArtists function with the mock server
	artists, err := FetchArtists(mockServer.URL + "/api/artists")
	if err != nil {
		t.Errorf("FetchArtists returned an error: %v", err)
		return
	}

	// Verify the returned artists
	if len(artists) != len(mockArtistsData) {
		t.Errorf("Expected %d artists, got %d", len(mockArtistsData), len(artists))
		return
	}

	for i, artist := range artists {
		mockArtist := mockArtistsData[i]
		if artist.ID != mockArtist.ID {
			t.Errorf("Expected artist ID %d, got %d", mockArtist.ID, artist.ID)
		}
		if artist.Name != mockArtist.Name {
			t.Errorf("Expected artist name %s, got %s", mockArtist.Name, artist.Name)
		}
		if artist.Image != mockArtist.Image {
			t.Errorf("Expected artist image %s, got %s", mockArtist.Image, artist.Image)
		}
		if artist.CreationDate != mockArtist.CreationDate {
			t.Errorf("Expected artist creation date %d, got %d", mockArtist.CreationDate, artist.CreationDate)
		}
		if artist.FirstAlbum != mockArtist.FirstAlbum {
			t.Errorf("Expected artist first album %s, got %s", mockArtist.FirstAlbum, artist.FirstAlbum)
		}
		if len(artist.Members) != len(mockArtist.Members) {
			t.Errorf("Expected %d members, got %d", len(mockArtist.Members), len(artist.Members))
			return
		}
		for j, member := range artist.Members {
			if member != mockArtist.Members[j] {
				t.Errorf("Expected member %s, got %s", mockArtist.Members[j], member)
			}
		}
	}
}
func TestFetchLocation(t *testing.T) {
	var mockLocationData []models.Location
	// Set up a mock server to handle the API request
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare the mock response data
		mockLocationData = []models.Location{
			{Locations: []string{"New York", "Los Angeles", "Chicago"}},
			{Locations: []string{"San Francisco", "Seattle", "Miami"}},
			{Locations: []string{"Boston", "Washington DC", "Atlanta"}},
		}

		// Encode the mock data as JSON
		mockLocationDataBytes, err := json.Marshal(mockLocationData)
		if err != nil {
			t.Errorf("Failed to marshal mock location data: %v", err)
			return
		}

		// Write the mock response
		w.Write(mockLocationDataBytes)
	}))
	defer mockServer.Close()

	// Call the FetchLocation function with the mock server
	locations, err := fetchLocation(mockServer.URL + "/api/locations")
	if err != nil {
		t.Errorf("FetchLocation returned an error: %v", err)
		return
	}

	// Verify the returned locations
	if len(locations) != len(mockLocationData) {
		t.Errorf("Expected %d locations, got %d", len(mockLocationData), len(locations))
		return
	}

	for i, location := range locations {
		mockLocation := mockLocationData[i]
		if len(location.Locations) != len(mockLocation.Locations) {
			t.Errorf("Expected %d locations, got %d", len(mockLocation.Locations), len(location.Locations))
			return
		}
		for j, loc := range location.Locations {
			if loc != mockLocation.Locations[j] {
				t.Errorf("Expected location %s, got %s", mockLocation.Locations[j], loc)
			}
		}
	}
}

func TestFetchDate(t *testing.T) {
	var mockDateData []models.Date
	// Set up a mock server to handle the API request
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare the mock response data
		mockDateData = []models.Date{
			{ID: 1, Dates: []string{"2023-01-01", "2023-01-02", "2023-01-03"}},
			{ID: 2, Dates: []string{"2023-02-01", "2023-02-02", "2023-02-03"}},
			{ID: 3, Dates: []string{"2023-03-01", "2023-03-02", "2023-03-03"}},
		}

		// Encode the mock data as JSON
		mockDateDataBytes, err := json.Marshal(mockDateData)
		if err != nil {
			t.Errorf("Failed to marshal mock date data: %v", err)
			return
		}

		// Write the mock response
		w.Write(mockDateDataBytes)
	}))
	defer mockServer.Close()

	// Call the FetchDate function with the mock server
	dates, err := fetchDate(mockServer.URL + "/api/dates")
	if err != nil {
		t.Errorf("FetchDate returned an error: %v", err)
		return
	}

	// Verify the returned dates
	if len(dates) != len(mockDateData) {
		t.Errorf("Expected %d dates, got %d", len(mockDateData), len(dates))
		return
	}

	for i, date := range dates {
		mockDate := mockDateData[i]
		if date.ID != mockDate.ID {
			t.Errorf("Expected date ID %d, got %d", mockDate.ID, date.ID)
		}
		if len(date.Dates) != len(mockDate.Dates) {
			t.Errorf("Expected %d dates, got %d", len(mockDate.Dates), len(date.Dates))
			return
		}
		for j, d := range date.Dates {
			if d != mockDate.Dates[j] {
				t.Errorf("Expected date %s, got %s", mockDate.Dates[j], d)
			}
		}
	}
}

func TestFetchRelationData(t *testing.T) {
	mockRelationData := []map[string][]string{}

	// Set up a mock server to handle the API request
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare the mock response data
		mockRelationData = []map[string][]string{
			{
				"2023-01-01": {"New York", "Los Angeles", "Chicago"},
				"2023-01-02": {"San Francisco", "Seattle", "Miami"},
				"2023-01-03": {"Boston", "Washington DC", "Atlanta"},
			},
			{
				"2023-02-01": {"New York", "Los Angeles", "Chicago"},
				"2023-02-02": {"San Francisco", "Seattle", "Miami"},
				"2023-02-03": {"Boston", "Washington DC", "Atlanta"},
			},
			{
				"2023-03-01": {"New York", "Los Angeles", "Chicago"},
				"2023-03-02": {"San Francisco", "Seattle", "Miami"},
				"2023-03-03": {"Boston", "Washington DC", "Atlanta"},
			},
		}

		// Encode the mock data as JSON
		mockRelationDataBytes, err := json.Marshal(mockRelationData)
		if err != nil {
			t.Errorf("Failed to marshal mock relation data: %v", err)
			return
		}

		// Write the mock response
		w.Write(mockRelationDataBytes)
	}))
	defer mockServer.Close()

	// Call the FetchRelationData function with the mock server
	relations, err := fetchRelationData(mockServer.URL + "/api/relation")
	if err != nil {
		t.Errorf("FetchRelationData returned an error: %v", err)
		return
	}

	// Verify the returned relations
	if len(relations) != len(mockRelationData) {
		t.Errorf("Expected 3 relations, got %d", len(relations))
		return
	}

	for i, relation := range relations {
		for date, locations := range relation.DatesLocations {
			if len(locations) != 3 {
				t.Errorf("Expected 3 locations for date %s in relation %d, got %d", date, i, len(locations))
				return
			}
		}
	}
}

func fetchLocation(url string) ([]models.Location, error) {
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

	var locations []models.Location
	err3 := json.Unmarshal(locationData, &locations)
	if err3 != nil {
		log.Println(err3)
		return []models.Location{}, err3
	}

	return locations, nil
}

func fetchDate(url string) ([]models.Date, error) {
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

	var dates []models.Date
	err3 := json.Unmarshal(datesData, &dates)
	if err3 != nil {
		log.Println(err3)
		return []models.Date{}, err3
	}

	return dates, nil
}

func fetchRelationData(url string) ([]models.Relation, error) {
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

	var relations []models.Relation
	err3 := json.Unmarshal(relationData, &relations)
	if err3 != nil {
		log.Println(err3)
		return []models.Relation{}, err3
	}

	return relations, nil
}
