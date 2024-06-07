package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bkmakwana2002/go-backend/usecase"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongoDB Connected")

}

func main() {
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := usecase.EmployeeService{MongoCollection: coll}

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	r.HandleFunc("/employee", empService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empService.GetEmployeeById).Methods(http.MethodGet)
	r.HandleFunc("/employee", empService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empService.UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empService.DeleteEmployee).Methods(http.MethodDelete)
	r.HandleFunc("/employee", empService.DeleteAll).Methods(http.MethodDelete)

	log.Println("Server running on port " + os.Getenv("PORT"))

	var str string = fmt.Sprintf(":%v", os.Getenv("PORT"))

	http.ListenAndServe(str, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running"))
}
