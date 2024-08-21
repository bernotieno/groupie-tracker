package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArtistInfo(t *testing.T) {
	// Backup the original log output
	originalLog := log.Writer()

	// Redirect log output to discard (no-op writer) to suppress log messages during testing
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(originalLog) // Restore original log output after test

	// Test cases
	testCases := []struct {
		name           string
		method         string
		path           string
		formValue      string
		expectedStatus int
	}{

		{
			name:           "Invalid method",
			method:         http.MethodGet,
			path:           "/artistInfo",
			formValue:      "Artist1",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid path",
			method:         http.MethodPost,
			path:           "/invalidPath",
			formValue:      "Artist1",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Internal Server Error from CollectData",
			method:         http.MethodPost,
			path:           "/artistInfo",
			formValue:      "Artist1",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request
			req := httptest.NewRequest(tc.method, tc.path, nil)
			if tc.formValue != "" {
				req.PostForm = map[string][]string{"ArtistName": {tc.formValue}}
			}

			// Create a new response recorder
			rr := httptest.NewRecorder()

			// Call the ArtistInfo function
			ArtistInfo(rr, req)

			// Verify the response status code
			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, rr.Code)
			}
		})
	}
}
