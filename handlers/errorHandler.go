package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"learn.zone01kisumu.ke/git/rcaleb/groupie-tracker/models"
)

func HandleError(w http.ResponseWriter, errMsg string, statusCode int) {

	w.WriteHeader(statusCode)
	// Parse the error template file
	tmp, err := template.ParseFiles("./templates/errorPage.html")
	if err != nil {
		// Log the parsing error
		log.Printf("Error parsing template: %v\n", err)
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
