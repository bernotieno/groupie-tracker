package handlers

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
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

// checkURL performs an HTTP GET request to the specified URL and sends an 
// error message to the channel if the request fails or returns a non-200 status code.
// If the request is successful with a 200 status code, it sends an empty string.
func checkURL(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		errorMsg := handleNetworkError(err)
		ch <- errorMsg
		return
	}

	if resp.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("unexpected status code %d for URL %s", resp.StatusCode, url)
		ch <- errorMsg
		return
	}

	ch <- ""
}

// handleNetworkError categorizes and returns a descriptive message for various network-related errors.
func handleNetworkError(err error) string {
	if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			return fmt.Sprintf("connection timed out: %v", err)
		}
		return fmt.Sprintf("network error: %v", err)
	}

	if os.IsNotExist(err) {
		return fmt.Sprintf("DNS resolution error: %v", err)
	}

	if os.IsPermission(err) {
		return fmt.Sprintf("connection refused: %v", err)
	}

	return "General network error"
}

// checkInternetConnection checks multiple URLs concurrently to determine if 
// the internet connection is active, returning an error message if any URL fails.
func checkInternetConnection() string {
	urls := []string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}

	// Channel to receive error messages
	ch := make(chan string, len(urls)) // Buffer the channel to avoid blocking

	var wg sync.WaitGroup
	wg.Add(len(urls))

	// Launch goroutines for each URL
	for _, url := range urls {
		go checkURL(url, ch, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(ch)

	// Collect results
	for msg := range ch {
		if msg != "" {
			return msg
		}
	}

	return ""
}
