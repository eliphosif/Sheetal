package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/eliphosif/Sheetal/model"
	"github.com/gorilla/mux"
)

type User struct {
	id   int
	name string
	age  int
}

type AnswerResponse struct {
	User
	answer string
}

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
	fmt.Println("inside digits forward")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the parameters
	//digitspan/digitforward/item/{itemid}/trail/{trailid}
	itemid, err := strconv.Atoi(params["itemid"])
	if err != nil {
		log.Fatal(err)
	}
	trailid, err := strconv.Atoi(params["trailid"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(itemid, trailid)

	itemsLength := len(model.Fmodal.Items)
	if itemid <= itemsLength {

		trailLength := len(model.Fmodal.Items[itemid].Trails)
		if trailid > trailLength {
			fmt.Fprintf(w, "please correct your url, Thankyou")
			return
		}
	} else {
		fmt.Fprintf(w, "please correct your url, Thankyou")
		return
	}
	if r.Method == "GET" {
		q := model.Fmodal.Items[itemid].Trails[trailid].Question
		fmt.Fprintf(w, q)
		return
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
