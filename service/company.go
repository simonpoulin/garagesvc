package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CompanyCreate ...
func CompanyCreate(payload model.CompanyCreatePayload) (companyID primitive.ObjectID, err error) {
	var company model.Company

	//Set data for new company
	company.ID = primitive.NewObjectID()
	company.Active = false
	company.Address = payload.Address
	company.Name = payload.Name
	company.Location = payload.Location

	//Insert to database
	err = dao.CompanyCreate(company)
	companyID = company.ID
	return
}

// CompanyDetail ...
func CompanyDetail(id primitive.ObjectID) (company model.Company, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for company from database
	company, err = dao.CompanyFindOne(filter)
	return
}

// CompanyList ...
func CompanyList(name string, page int) (companyList interface{}, err error) {
	var (
		filter = bson.M{}
	)

	if name != "" {
		filter = bson.M{"name": bson.M{"$regex": name}}
	}

	//Get companies
	companies, err := dao.CompanyFind(filter)

	//Paging list
	if page > 0 {
		companyList, err = util.Paging(companies, page, 8)
		return
	}
	companyList = companies

	return
}

// CompanyUpdate ...
func CompanyUpdate(id primitive.ObjectID, payload model.CompanyUpdatePayload, active string) (companyID primitive.ObjectID, err error) {

	var update bson.M

	//Set filter and data
	filter := bson.M{"_id": id}

	if active != "" {
		stt, _ := strconv.ParseBool(active)
		update = bson.M{"$set": bson.M{"active": stt}}
	} else {
		update = bson.M{"$set": bson.M{
			"active":   payload.Active,
			"address":  payload.Address,
			"name":     payload.Name,
			"location": payload.Location,
		}}
	}

	//Update company
	err = dao.CompanyUpdateOne(filter, update)

	//Return data
	companyID = id
	return
}

// CompanyDelete ...
func CompanyDelete(id primitive.ObjectID) (err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Delete company
	err = dao.CompanyDelete(filter)
	return
}
