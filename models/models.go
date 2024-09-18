package models

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Data struct {
	A Artist
	R Relation
	L Location
	D Date
}

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
}

type Location struct {
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Err struct {
	ErrMsg     string
	StatusCode int
}

type SearchIndex struct {
	ArtistName   map[string][]IndexedData
	MemberName   map[string][]IndexedData
	LocationName map[string][]IndexedData
	FirstAlbum   map[string][]IndexedData
	CreationDate map[int][]IndexedData
	mu           sync.RWMutex
}

type IndexedData struct {
	Data       Data
	ArtistName string
}

type SearchResult struct {
	Result     string
	ArtistName string
}

func (index *SearchIndex) Search(query string) []SearchResult {
	query = strings.ToLower(query)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultChan := make(chan SearchResult, 100)

	var wg sync.WaitGroup

	searchFunctions := []func(context.Context, string, chan<- SearchResult){
		index.searchArtistNames,
		index.searchMemberNames,
		index.searchLocations,
		index.searchFirstAlbums,
		index.searchCreationDates,
	}

	for _, searchFunc := range searchFunctions {
		wg.Add(1)
		go func(sf func(context.Context, string, chan<- SearchResult)) {
			defer wg.Done()
			sf(ctx, query, resultChan)
		}(searchFunc)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []SearchResult
	for result := range resultChan {
		select {
		case <-ctx.Done():
			return results
		default:
			results = append(results, result)
		}
	}

	return results
}

func (index *SearchIndex) searchArtistNames(ctx context.Context, query string, resultChan chan<- SearchResult) {
	index.mu.RLock()
	defer index.mu.RUnlock()

	for name, data := range index.ArtistName {
		if strings.Contains(strings.ToLower(name), query) {
			for _, d := range data {
				select {
				case <-ctx.Done():
					return
				case resultChan <- SearchResult{
					Result:     d.ArtistName + " - artist/band",
					ArtistName: d.ArtistName,
				}:
				}
			}
		}
	}
}

func (index *SearchIndex) searchMemberNames(ctx context.Context, query string, resultChan chan<- SearchResult) {
	index.mu.RLock()
	defer index.mu.RUnlock()
	for _, data := range index.MemberName {
		for _, d := range data {
			for _, member := range d.Data.A.Members {
				if strings.Contains(strings.ToLower(member), query) {
					select {
					case <-ctx.Done():
						return
					case resultChan <- SearchResult{
						Result:     member + " - member of " + d.ArtistName + " band",
						ArtistName: d.ArtistName,
					}:
					}
					break
				}
			}
		}
	}
}

func (index *SearchIndex) searchLocations(ctx context.Context, query string, resultChan chan<- SearchResult) {
	index.mu.RLock()
	defer index.mu.RUnlock()
	for _, data := range index.LocationName {
		for _, d := range data {
			for _, location := range d.Data.L.Locations {
				if strings.Contains(strings.ToLower(location), query) {
					select {
					case <-ctx.Done():
						return
					case resultChan <- SearchResult{
						Result:     location + " - location-" + d.ArtistName,
						ArtistName: d.ArtistName,
					}:
					}
					break
				}
			}
		}
	}
}

func (index *SearchIndex) searchFirstAlbums(ctx context.Context, query string, resultChan chan<- SearchResult) {
	index.mu.RLock()
	defer index.mu.RUnlock()
	for _, data := range index.FirstAlbum {
		for _, d := range data {
			if strings.Contains(strings.ToLower(d.Data.A.FirstAlbum), query) {
				select {
				case <-ctx.Done():
					return
				case resultChan <- SearchResult{
					Result:     d.Data.A.FirstAlbum + " - first album-" + d.ArtistName,
					ArtistName: d.ArtistName,
				}:
				}
				break
			}
		}
	}
}

func (index *SearchIndex) searchCreationDates(ctx context.Context, query string, resultChan chan<- SearchResult) {
	if creationDate, err := strconv.Atoi(query); err == nil {
		index.mu.RLock()
		data, found := index.CreationDate[creationDate]
		index.mu.RUnlock()

		if found {
			for _, d := range data {
				select {
				case <-ctx.Done():
					return
				case resultChan <- SearchResult{
					Result:     strconv.Itoa(d.Data.A.CreationDate) + " - creation date-" + d.ArtistName,
					ArtistName: d.ArtistName,
				}:
				}
			}
		}
	}
}

// PreloadData takes the preloaded data and populates the SearchIndex.
func (index *SearchIndex) PreloadData(data []Data) {
	index.mu.Lock()
	defer index.mu.Unlock()

	for _, item := range data {
		// Index artist names
		name := strings.ToLower(item.A.Name)
		index.ArtistName[name] = append(index.ArtistName[name], IndexedData{
			Data:       item,
			ArtistName: item.A.Name,
		})

		// Index member names
		for _, member := range item.A.Members {
			memberLower := strings.ToLower(member)
			index.MemberName[memberLower] = append(index.MemberName[memberLower], IndexedData{
				Data:       item,
				ArtistName: item.A.Name,
			})
		}

		// Index locations
		for _, location := range item.L.Locations {
			locationLower := strings.ToLower(location)
			index.LocationName[locationLower] = append(index.LocationName[locationLower], IndexedData{
				Data:       item,
				ArtistName: item.A.Name,
			})
		}

		// Index first albums
		albumLower := strings.ToLower(item.A.FirstAlbum)
		index.FirstAlbum[albumLower] = append(index.FirstAlbum[albumLower], IndexedData{
			Data:       item,
			ArtistName: item.A.Name,
		})

		// Index creation dates
		index.CreationDate[item.A.CreationDate] = append(index.CreationDate[item.A.CreationDate], IndexedData{
			Data:       item,
			ArtistName: item.A.Name,
		})
	}
}

// SearchHandler handles search requests using preloaded data
func (index *SearchIndex) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Get the search query from URL parameters
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing search query parameter 'q'", http.StatusBadRequest)
		return
	}

	// Perform the search on the preloaded data
	results := index.Search(query)

	// Send search results back as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
