package repository

import (
	"context"

	"github.com/Bkmakwana2002/go-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeBYId(empId string) (*model.Employee, error) {
	var emp model.Employee
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "employee_id", Value: empId}}).Decode(&emp)

	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployees() ([]model.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []model.Employee
	err = results.All(context.Background(), &emps)
	if err != nil {
		return nil, err
	}
	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployee(empId string, updateEmp *model.Employee) (int64, error) {
	results, err := r.MongoCollection.UpdateByID(context.Background(), bson.D{
		{
			Key:   "employee_id",
			Value: empId,
		},
	}, bson.D{{
		Key:   "$set",
		Value: updateEmp,
	}})

	if err != nil {
		return 0, err
	}
	return results.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployee(empId string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{
		Key:   "employee_id",
		Value: empId,
	}})

	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmp() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
