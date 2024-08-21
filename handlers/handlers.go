package handlers

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
)

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
