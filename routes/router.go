package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/eliphosif/Sheetal/model"
	"github.com/gorilla/mux"
)

type User struct {
	Id    int    `json:"userid"`
	Email string `json:"useremail"`
	Name  string `json:"username"`
	Age   int    `json:"userage"`
}

var limit bool = false

type AnswerReq struct {
	Userid int
	answer string
}
type QuestionRes struct {
	Userid   int
	Question string
}

type AnswerRes struct {
	User
	limit bool
}
type UserRegistrationError struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

var userRegistrationError UserRegistrationError = UserRegistrationError{
	Message:     "Not registered",
	Description: "it seems you're not registered user, Please register ang try again ! :)",
}

var NewUser = User{}

func isRegistered(r *http.Request) bool {
	if NewUser.Id == 0 {
		return true
	}
	return false
}
func UserRegister(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&NewUser)

	NewUser.Id = int(time.Now().UnixNano()) // current local time
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewUser)
}

func DigitsForward(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {

		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return

	}

	fmt.Println(r.Body)
	w.Header().Set("Content-Type", "application/json")

	if isRegistered(r) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(userRegistrationError)
		return
	}
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
	q := model.Fmodal.Items[itemid].Trails[trailid].Question
	if r.Method == "GET" {
		quesreq := QuestionRes{
			Userid:   NewUser.Id,
			Question: q,
		}
		json.NewEncoder(w).Encode(quesreq)
		return
	} else if r.Method == "POST" {
		var ansRes AnswerRes
		var ansReq AnswerReq
		err := json.NewDecoder(r.Body).Decode(&ansReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(ansReq.answer, q, r.Body)
		if ansReq.answer == q {
			ansRes.User.Id = NewUser.Id
			ansRes.limit = false

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(ansRes)
		} else {

		}
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
