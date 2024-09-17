package handlers

import (
	"fmt"
	"net/http"
	"os"
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

func checkInternetConnection() string {
	urls := []string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}

	errCh := make(chan string, len(urls)) // Channel to receive errors
	doneCh := make(chan bool)             // Channel to signal completion

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				errCh <- fmt.Sprintf("Error connecting to %s: %v", url, err)
				return
			}
			errCh <- "" // No error
		}(url)
	}

	// Wait for all checks to complete
	go func() {
		for i := 0; i < len(urls); i++ {
			errMsg := <-errCh
			if errMsg != "" {
				doneCh <- false // Error occurred
				return
			}
		}
		doneCh <- true // All checks successful
	}()

	if success := <-doneCh; !success {
		return <-errCh
	}

	return ""
}
