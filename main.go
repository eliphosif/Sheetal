package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/eliphosif/Sheetal/routes"
	"github.com/eliphosif/Sheetal/user"

	"github.com/gorilla/mux"
)

var myUser user.User

func main() {

	initlizeRouter()
}

func initlizeRouter() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	r := mux.NewRouter()

	r.HandleFunc("/home", home)
	r.HandleFunc("/register", routes.UserRegister)
	r.HandleFunc("/digitspan/digitforward/item/{itemid}/trail/{trailid}", routes.DigitsForward)
	r.HandleFunc("/digitspan/digitbackward/item/{itemid}/trail/{trailid}", routes.DigitsBackward)
	r.HandleFunc("/digitspan/LetterNumberSequencing/item/{itemid}/trail/{trailid}", routes.LetterNumberSequencing)

	fmt.Println("server is listening:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home")
}

/*

	  = initlizeMongoConnection()
var Survey *mongo.Collection
var ctx = context.Background()
func initlizeMongoConnection() (result []*mongo.Collection) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")

	MongoDBURI := "mongodb+srv://" + USERNAME + ":" + PASSWORD + "@cluster0.0imgv.mongodb.net/test?authSource=admin&replicaSet=atlas-fydu9m-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"

	//defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MongoDBURI))

	golangMongoDB := client.Database("Survey")
	DigitsForward := golangMongoDB.Collection("DigitsForward")
	result = (append(result, DigitsForward))
	DigitsBackward := golangMongoDB.Collection("DigitsBackward")
	result = (append(result, DigitsBackward))
	LetterNumberSequencing := golangMongoDB.Collection("LetterNumberSequencing")
	result = (append(result, LetterNumberSequencing))

	return result

}
*/
