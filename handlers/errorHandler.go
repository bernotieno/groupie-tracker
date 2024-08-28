package handlers

import (
	"fmt"
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

	w.WriteHeader(statusCode)
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

	err1 := tmp.Execute(w, errors)
	if err1 != nil {
		
		fmt.Printf("Error executing template here: %v\n", err1)
		
		if statusCode != http.StatusInternalServerError {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
}
