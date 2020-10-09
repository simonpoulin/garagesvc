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
func EmployeeDetail(id primitive.ObjectID) (employeeRes model.EmployeeResponse, err error) {

	filter := bson.M{"_id": id}

	//Looking for employee from database
	employee, err := dao.EmployeeFindOne(filter)
	if err != nil {
		return
	}

	employeeRes, err = EmployeeConvertToResponse(employee)

	return
}

// EmployeeList ...
func EmployeeList(query model.AppQuery) (employeeList util.PagedList, err error) {

	var (
		findQuery       = query.GenerateFindQuery()
		employeeListRes []model.EmployeeResponse
	)

	//Get employee list
	employees, err := dao.EmployeeFind(findQuery)
	if err != nil {
		return
	}

	//Get employee response list
	for _, employee := range employees {
		var employeeRes model.EmployeeResponse
		employeeRes, err = EmployeeConvertToResponse(employee)
		if err != nil {
			return
		}
		employeeListRes = append(employeeListRes, employeeRes)
	}

	//Paging list
	employeeList, err = util.Paging(employeeListRes, query.Page, 8)

	return
}

// EmployeeUpdate ...
func EmployeeUpdate(id primitive.ObjectID, payload model.EmployeeUpdatePayload, active string) (employeeID primitive.ObjectID, err error) {

	var update bson.M

	//Set filter and data
	filter := bson.M{"_id": id}

	//Get old employee info
	employee, err := dao.EmployeeFindOne(filter)
	if err != nil {
		return
	}

	//Check payload password
	if employee.Password != payload.Password {
		payload.Password = util.Hash(payload.Password)
	}

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

// EmployeeConvertToResponse ...
func EmployeeConvertToResponse(e model.Employee) (res model.EmployeeResponse, err error) {

	res.ID = e.ID
	res.Name = e.Name
	res.Phone = e.Phone
	res.Password = e.Password
	res.Active = e.Active

	return
}
