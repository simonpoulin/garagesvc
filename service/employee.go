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
func EmployeeList(active string, name string, page int) (employeeList interface{}, err error) {
	var (
		filterParts []bson.M
		findQuery   []bson.M
	)

	//Set filter parts
	if active != "" {
		stt, _ := strconv.ParseBool(active)
		filterParts = append(filterParts, bson.M{"active": stt})
	}

	if name != "" {
		filterParts = append(filterParts, bson.M{"name": bson.M{"$regex": name}})
	}

	//Set filter query from parts
	findQuery = append(findQuery, bson.M{"$match": func() bson.M {
		if filterParts != nil {
			if len(filterParts) > 0 {
				return bson.M{"$and": filterParts}
			}
		}
		return bson.M{}
	}()})

	//Get employee list
	employees, err := dao.EmployeeFind(findQuery)

	//Paging list
	if page > 0 {
		employeeList, err = util.Paging(employees, page, 8)
		return
	}
	employeeList = employees

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
		update = bson.M{"$set": bson.M{
			"active":   payload.Active,
			"password": util.Hash(payload.Password),
			"name":     payload.Name,
			"phone":    payload.Phone,
		}}
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
