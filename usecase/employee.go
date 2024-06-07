package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Bkmakwana2002/go-backend/model"
	"github.com/Bkmakwana2002/go-backend/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	insertId, err := repo.InsertEmployee(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insert Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID

	w.WriteHeader(http.StatusOK)

	log.Println("Employee Inserted", insertId)
}

func (svc *EmployeeService) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empId := mux.Vars(r)["id"]
	log.Println(empId)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeBYId(empId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Found", emp)

}

func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployees()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Found", emp)
}

func (svc *EmployeeService) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	empId := mux.Vars(r)["id"]

	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request")
		res.Error = "Empty Emp Id"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empId
	count, err := repo.UpdateEmployee(empId, &emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Updated", count)

}

func (svc *EmployeeService) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empId := mux.Vars(r)["id"]
	log.Println(empId)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteEmployee(empId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Deleted", count)
}

func (svc *EmployeeService) DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteAllEmp()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Deleted", count)
}
