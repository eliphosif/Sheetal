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

		fmt.Fprintf(w, "please register to take the survey, Thankyou")
		return
	}
}
func DigitsForward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	} else {

		fmt.Fprintf(w, "please register to take the survey, Thankyou")
		return
	}
}
func DigitsBackward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	} else {

		fmt.Fprintf(w, "please register to take the survey, Thankyou")
		return
	}
}
func LetterNumberSequencing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	} else {

		fmt.Fprintf(w, "please register to take the survey, Thankyou")
		return
	}
}

/*
DigitsForward
DigitsBackward
LetterNumberSequencing
*/
