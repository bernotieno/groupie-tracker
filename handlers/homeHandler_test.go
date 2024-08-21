package handlers

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	originalLog := log.Writer()

	// Redirect log output to discard (no-op writer) to suppress log messages during testing
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(originalLog) // Restore original log output after test

	// Test cases
	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "Invalid method",
			method:         http.MethodPost,
			path:           "/",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid path",
			method:         http.MethodGet,
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Error parsing template",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Error fetching artists",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request
			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
				return
			}

			// Create a new response recorder
			rr := httptest.NewRecorder()

			// Call the Home function
			Home(rr, req)

			// Verify the response status code
			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, rr.Code)
			}
		})
	}
}

// func HandleError(w http.ResponseWriter, message string, statusCode int) {
// 	http.Error(w, message, statusCode)
// }

// func checkInternetConnection() string {
// 	// Simulate internet connectivity issues
// 	if strings.Contains(t.Name(), "Internet connectivity issues") {
// 		return "Internet connectivity issues"
// 	}
// 	return ""
// }

// func FetchArtists() (interface{}, error) {
// 	// Simulate error fetching artists
// 	if strings.Contains(t.Name(), "Error fetching artists") {
// 		return nil, ErrFetchingArtists
// 	}
// 	return []interface{}{}, nil
// }

var ErrFetchingArtists = errors.New("error fetching artists")
