package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/api"
	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if strings.ToUpper(r.Method) != "GET" {
		HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./templates/Home.html")
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	msg := checkInternetConnection()
	if msg != "" {
		fmt.Println(msg)
		HandleError(w, "Internet Connectivity Issues", http.StatusRequestTimeout)
		return
	}
	data, err := api.FetchArtists()
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err1 := tmp.Execute(w, data)
	if err1 != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

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

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		HandleError(w, "400 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/artistInfo" {
		HandleError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("./templates/artistPage.html")
	if err != nil {
		HandleError(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	artistName := r.FormValue("ArtistName")

	if artistName == "" {
		HandleError(w, "Bad Request: Missing artist name", http.StatusBadRequest)
		return
	}
	msg := checkInternetConnection()
	if msg != "" {
		fmt.Println(msg)
		HandleError(w, "Poor Internet Connectivity", http.StatusRequestTimeout)
		return
	}

	artistInfo, err := api.CollectData()
	if err != nil {
		HandleError(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	var art models.Data
	found := false
	for _, artist := range artistInfo {
		if artistName == artist.A.Name {
			art = artist
			found = true
			break
		}
	}

	if !found {
		HandleError(w, "404 Artist Not Found", http.StatusNotFound)
		return
	}

	if err := t.Execute(w, art); err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}

func HandleError(w http.ResponseWriter, errMsg string, statusCode int) {

	w.WriteHeader(statusCode)
	// Parse the error template file
	tmp, err := template.ParseFiles("./templates/errorPage.html")
	if err != nil {
		// Log the parsing error
		fmt.Printf("Error parsing template: %v\n", err)
		// Ensure that no further status code is set by returning immediately
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create an error message object
	errors := models.Err{
		ErrMsg:     errMsg,
		StatusCode: statusCode,
	}

	// Execute the template
	err1 := tmp.Execute(w, errors)
	if err1 != nil {
		// Log the template execution error and respond with 500 if not already handled
		fmt.Printf("Error executing template here: %v\n", err1)
		// Ensure that we do not set another status code by handling the error internally
		if statusCode != http.StatusInternalServerError {
			// Only set the internal server error if no other status code was set
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
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
