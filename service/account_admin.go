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
	var employee model.EmployeeCreateBSON

	//Set data for new employee
	employee = payload.ConvertToCreateBSON()

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
	if err != nil {
		return
	}

	//Check password match
	if payload.Password != employee.Password {
		err = errors.New("password not match")
		return
	}

	//Check if account is active
	if !employee.Active {
		err = errors.New("employee disabled")
		return
	}

	//Generate token
	token, err = employee.GenerateToken()
	return
}
