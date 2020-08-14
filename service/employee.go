package service

import (
	"context"
	"errors"
	"garagesvc/model"
	"garagesvc/module/mongodb"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeCreate ...
func EmployeeCreate(payload model.EmployeeCreatePayload) (employeeID string, err error) {
	var (
		employee    model.Employee
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Set data for new employee
	employee.ID = primitive.NewObjectID()
	employee.Active = true
	employee.Password = util.Hash(payload.Password)
	employee.Name = payload.Name
	employee.Phone = payload.Phone

	//Insert to database
	_, err = employeeCol.InsertOne(ctxt, employee)
	employeeID = employee.ID.Hex()
	return
}

// EmployeeLogin ...
func EmployeeLogin(payload model.EmployeeLoginPayload) (token string, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Get employee by phone number
	payload.Password = util.Hash(payload.Password)
	var e model.Employee
	filter := bson.M{"phone": payload.Phone}
	err = employeeCol.FindOne(ctxt, filter).Decode(&e)
	if err != nil {
		return
	}

	//Check password match
	if payload.Password != e.Password {
		err = errors.New("password not match")
		return
	}
	token, err = util.TokenEncode(e.ID.Hex())
	return
}

// EmployeeDetail ...
func EmployeeDetail(id string) (e model.Employee, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Looking for employee from database
	err = employeeCol.FindOne(ctxt, filter).Decode(&e)
	return
}

// EmployeeList ...
func EmployeeList() (employeeList []model.Employee, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Get employees
	cur, err := employeeCol.Find(ctxt, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add employees to list
	for cur.Next(ctxt) {
		var result model.Employee
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		employeeList = append(employeeList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// EmployeeListByActiveState ...
func EmployeeListByActiveState(active string) (employeeList []model.Employee, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Set filter
	filter := bson.M{"active": active}

	//Get employees

	cur, err := employeeCol.Find(ctxt, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add employees to list
	for cur.Next(ctxt) {
		var result model.Employee
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		employeeList = append(employeeList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// EmployeeUpdate ...
func EmployeeUpdate(id string, payload model.EmployeeUpdatePayload) (employeeID string, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Set filter and data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"active":   payload.Active,
		"password": payload.Password,
		"name":     payload.Name,
		"phone":    payload.Phone,
	}}

	//Update employee
	_, err = employeeCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	employeeID = id
	return
}

// EmployeeChangeActive ...
func EmployeeChangeActive(id string) (employeeStatus bool, err error) {
	var (
		employee    model.Employee
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Set filter active state data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"active": !employee.Active}}

	//Update employee
	_, err = employeeCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	employeeStatus = !employee.Active
	return
}

// EmployeeDelete ...
func EmployeeDelete(id string) (err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Delete employee
	_, err = employeeCol.DeleteOne(ctxt, filter)
	return
}
