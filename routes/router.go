package routes

import (
	"fmt"
	"net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	} else {

		fmt.Fprintf(w, "Please provide valid login details")
		return
	}
}
func DigitsBackward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	} else {

		fmt.Fprintf(w, "Please provide valid login details")
		return
	}
}
func LetterNumberSequencing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	} else {

		fmt.Fprintf(w, "Please provide valid login details")
		return
	}
}

/*
DigitsForward
DigitsBackward
LetterNumberSequencing
*/
