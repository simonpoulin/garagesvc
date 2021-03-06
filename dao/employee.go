package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// EmployeeCreate ...
func EmployeeCreate(employee model.EmployeeCreateBSON) (err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctx         = context.Background()
	)
	_, err = employeeCol.InsertOne(ctx, employee)
	return
}

// EmployeeFindOne ...
func EmployeeFindOne(filter bson.M) (employee model.Employee, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctx         = context.Background()
	)
	err = employeeCol.FindOne(ctx, filter).Decode(&employee)
	return
}

// EmployeeFind ...
func EmployeeFind(filter []bson.M) (employeeList []model.Employee, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctx         = context.Background()
	)
	// Looking for employees
	cur, err := employeeCol.Aggregate(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &employeeList)
	return
}

// EmployeeUpdateOne ...
func EmployeeUpdateOne(filter bson.M, data bson.M) (err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctx         = context.Background()
	)
	_, err = employeeCol.UpdateOne(ctx, filter, data)
	return
}

// EmployeeDelete ...
func EmployeeDelete(filter bson.M) (err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctx         = context.Background()
	)
	_, err = employeeCol.DeleteOne(ctx, filter)
	return
}
