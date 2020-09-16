package service

import (
	"errors"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeRegister ...
func EmployeeRegister(payload model.EmployeeRegisterPayload) (employeeID primitive.ObjectID, err error) {
	var employee model.Employee

	//Set data for new employee
	employee.ID = primitive.NewObjectID()
	employee.Active = true
	employee.Password = util.Hash(payload.Password)
	employee.Name = payload.Name
	employee.Phone = payload.Phone

	//Insert to database
	err = dao.EmployeeCreate(employee)
	employeeID = employee.ID
	return
}

// EmployeeLogin ...
func EmployeeLogin(payload model.EmployeeLoginPayload) (token string, err error) {

	//Get employee by phone number
	payload.Password = util.Hash(payload.Password)
	filter := bson.M{"phone": payload.Phone}
	employee, err := dao.EmployeeFindOne(filter)

	//Check password match
	if payload.Password != employee.Password {
		err = errors.New("password not match")
		return
	}
	if !employee.Active {
		err = errors.New("employee disabled")
		return
	}
	token, err = employee.GenerateToken()
	return
}
