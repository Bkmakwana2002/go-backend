package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Bkmakwana2002/go-backend/model"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/internal/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	mongoTestClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongoDB Connected")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongoDB pinged")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	//dummy data
	emp1 := "uuid.New().String()"
	// emp2 := uuid.New().String()

	// connect to collection
	coll := mongoTestClient.Database("companyDB").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// InsertEmployee
	t.Run("Insert Employee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "John",
			Department: "IT",
			EmployeeID: emp1,
		}

		t.Log(emp)

		res, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("------------------------", err)
		}
		t.Log(res)
	})

	// Getting Emp1
	t.Run("Get Employee", func(t *testing.T) {
		res, err := empRepo.FindEmployeeBYId(emp1)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(res)
	})

	//Get All Emps
	t.Run("Get All Employee", func(t *testing.T) {
		res, err := empRepo.FindAllEmployees()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	})

	// update an Emp
	t.Run("Update Employee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "John2",
			Department: "IT",
			EmployeeID: emp1,
		}
		res, err := empRepo.UpdateEmployee(emp1, &emp)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	})

	// Delete an Emp
	t.Run("Delete Employee", func(t *testing.T) {
		res, err := empRepo.DeleteEmployee(emp1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	})
}
