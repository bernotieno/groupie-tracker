package handlers

import (
	"html/template"
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

// HandleError writes an HTTP error response with a specified status code and message.
// It renders an error page using a template, logging errors encountered during
// template parsing or execution. If template execution fails, it ensures that
// a 500 Internal Server Error is returned if it was not already set.
func HandleError(w http.ResponseWriter, errMsg string, statusCode int) {
	// Set the status code using the `Header()` method before writing the body
	w.Header().Set("Content-Type", "text/html") // Set content type (optional, but good practice)
	w.WriteHeader(statusCode)
	// Parse the template file
	tmp, err := template.ParseFiles("./templates/errorPage.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create an error message object
	errors := models.Err{
		ErrMsg:     errMsg,
		StatusCode: statusCode,
	}

	// Execute the template and write it directly to the response
	if err := tmp.Execute(w, errors); err != nil {
		log.Printf("Error executing template: %v\n", err)
		// Fallback to a simple error message if template execution fails
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}

}
