package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eliphosif/Sheetal/user"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var myUser user.User

var AllCustomers *mongo.Collection

func main() {

	AllCustomers = initlizeMongoConnection()
	initlizeRouter()
}

func initlizeRouter() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", UserRegister).Methods("GET")
	r.HandleFunc("/digitspan/digitforward/item/{itemid}/trail/{trailid}", DigitsForward).Methods("GET")

	r.HandleFunc("/digitspan/digitbackward/item/{itemid}/trail/{trailid}", DigitsBackward).Methods("GET")
	r.HandleFunc("/api/customers", getCustomers).Methods("GET")
	r.HandleFunc("/api/customer/{id}", getCustomer).Methods("GET")

	fmt.Println("server is listening:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

func initlizeMongoConnection() *mongo.Collection {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")

	MongoDBURI := "mongodb+srv://" + USERNAME + ":" + PASSWORD + "@cluster0.0imgv.mongodb.net/test?authSource=admin&replicaSet=atlas-fydu9m-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"

	//defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MongoDBURI))

	golangMongoDB := client.Database("GolangMongo")
	AllCustomers := golangMongoDB.Collection("AllCustomers")
	return AllCustomers

}
