package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeDetail ...
func EmployeeDetail(id primitive.ObjectID) (employee model.Employee, err error) {

	filter := bson.M{"_id": id}

	//Looking for employee from database
	employee, err = dao.EmployeeFindOne(filter)
	return
}

// EmployeeList ...
func EmployeeList(query model.AppQuery) (employeeList util.PagedList, err error) {

	var findQuery = query.GenerateFindQuery()

	//Get employee list
	employees, err := dao.EmployeeFind(findQuery)
	if err != nil {
		return
	}

	//Paging list
	employeeList, err = util.Paging(employees, query.Page, 8)

	return
}

// EmployeeUpdate ...
func EmployeeUpdate(id primitive.ObjectID, payload model.EmployeeUpdatePayload, active string) (employeeID primitive.ObjectID, err error) {

	var update bson.M

	//Set filter and data
	filter := bson.M{"_id": id}

	if active != "" {
		stt, _ := strconv.ParseBool(active)
		update = bson.M{"$set": bson.M{"active": stt}}
	} else {
		update = bson.M{"$set": payload.ConvertToUpdateBSON()}
	}

	//Update employee
	err = dao.EmployeeUpdateOne(filter, update)

	//Return data
	employeeID = id
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
