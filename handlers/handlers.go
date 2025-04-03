package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// StaticServer serves static files from the server's root directory based on
// the request URL path. It returns a 500 Internal Server Error if the file
// cannot be accessed or 403 Forbidden if the requested path is a directory.
// Otherwise, it serves the requested file.
func StaticServer(w http.ResponseWriter, r *http.Request) {
	filePath := "." + r.URL.Path
	info, err := os.Stat(filePath)
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if info.IsDir() {
		HandleError(w, "Acces Forbiden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filePath)
}

// checkInternetConnection attempts to make an HTTP GET request to a reliable server.
// It returns an error if the request fails, indicating no internet connection or instability.
func CheckInternetConnection() error {
	client := http.Client{
		Timeout: 5 * time.Second, // Set a timeout to avoid hanging requests
	}

	// Try making a GET request to a reliable server (e.g., google.com)
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		return fmt.Errorf("internet connectivity issue: %v", err)
	}
	defer resp.Body.Close()

	// Check if the server responded with a valid status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("internet connectivity issue: received status code %d", resp.StatusCode)
	}

	return nil
}
