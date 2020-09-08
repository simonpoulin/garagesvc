package service

import (
	"garagesvc/dao"
	"garagesvc/model"

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
func CompanyList() (companyList []model.Company, err error) {

	//Get companys
	companyList, err = dao.CompanyFind(bson.M{})
	return
}

// CompanyUpdate ...
func CompanyUpdate(id primitive.ObjectID, payload model.CompanyUpdatePayload) (companyID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"active":   payload.Active,
		"address":  payload.Address,
		"name":     payload.Name,
		"location": payload.Location,
	}}

	//Update company
	err = dao.CompanyUpdateOne(filter, update)

	//Return data
	companyID = id
	return
}

// CompanyChangeActive ...
func CompanyChangeActive(id primitive.ObjectID) (companyID primitive.ObjectID, err error) {

	//Set filter
	filter := bson.M{"_id": id}
	company, err := dao.CompanyFindOne(filter)
	if err != nil {
		return
	}

	//Set active state data
	update := bson.M{"$set": bson.M{"active": !company.Active}}

	//Update company
	err = dao.CompanyUpdateOne(filter, update)

	//Return data
	companyID = company.ID
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
