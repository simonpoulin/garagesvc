package service

import (
	"errors"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeCreate ...
func EmployeeCreate(payload model.EmployeeCreatePayload) (employeeID primitive.ObjectID, err error) {
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
	token, err = employee.GenerateToken()
	return
}

// EmployeeDetail ...
func EmployeeDetail(id primitive.ObjectID) (employee model.Employee, err error) {

	filter := bson.M{"_id": id}

	//Looking for employee from database
	employee, err = dao.EmployeeFindOne(filter)
	return
}

// EmployeeList ...
func EmployeeList() (employeeList []model.Employee, err error) {

	//Get employees
	employeeList, err = dao.EmployeeFind(bson.M{})
	return
}

// EmployeeListByActiveState ...
func EmployeeListByActiveState(active string) (employeeList []model.Employee, err error) {

	//Set filter
	filter := bson.M{"active": active}

	//Get employees
	employeeList, err = dao.EmployeeFind(filter)
	return
}

// EmployeeUpdate ...
func EmployeeUpdate(id primitive.ObjectID, payload model.EmployeeUpdatePayload) (employeeID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"active":   payload.Active,
		"password": payload.Password,
		"name":     payload.Name,
		"phone":    payload.Phone,
	}}

	//Update employee
	err = dao.EmployeeUpdateOne(filter, update)

	//Return data
	employeeID = id
	return
}

// EmployeeChangeActive ...
func EmployeeChangeActive(id primitive.ObjectID) (employeeID primitive.ObjectID, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Get employee
	employee, err := dao.EmployeeFindOne(filter)
	if err != nil {
		return
	}

	//Set active state data
	update := bson.M{"$set": bson.M{"active": !employee.Active}}

	//Update employee
	err = dao.EmployeeUpdateOne(filter, update)

	//Return data
	employeeID = employee.ID
	return
}

// EmployeeDelete ...
func EmployeeDelete(id primitive.ObjectID) (err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Delete employee
	err = dao.EmployeeDelete(filter)
	return
}
