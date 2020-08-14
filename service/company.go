package service

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CompanyCreate ...
func CompanyCreate(payload model.CompanyCreatePayload) (companyID string, err error) {
	var (
		company    model.Company
		companyCol = mongodb.CompanyCol()
		ctxt       = context.Background()
	)

	//Set data for new company
	company.ID = primitive.NewObjectID()
	company.Active = false
	company.Address = payload.Address
	company.Name = payload.Name
	company.Location = payload.Location

	//Insert to database
	_, err = companyCol.InsertOne(ctxt, company)
	companyID = company.ID.Hex()
	return
}

// CompanyDetail ...
func CompanyDetail(id string) (e model.Company, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctxt       = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Looking for company from database
	err = companyCol.FindOne(ctxt, filter).Decode(&e)
	return
}

// CompanyList ...
func CompanyList() (companyList []model.Company, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctxt       = context.Background()
	)

	//Get companys
	cur, err := companyCol.Find(ctxt, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add companys to list
	for cur.Next(ctxt) {
		var result model.Company
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		companyList = append(companyList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// CompanyUpdate ...
func CompanyUpdate(id string, payload model.CompanyUpdatePayload) (companyID string, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctxt       = context.Background()
	)

	//Set filter and data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"active":   payload.Active,
		"address":  payload.Address,
		"name":     payload.Name,
		"location": payload.Location,
	}}

	//Update company
	_, err = companyCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	companyID = id
	return
}

// CompanyChangeActive ...
func CompanyChangeActive(id string) (companyStatus bool, err error) {
	var (
		company    model.Company
		companyCol = mongodb.CompanyCol()
		ctxt       = context.Background()
	)

	//Set filter and active state data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"active": !company.Active}}

	//Update company
	_, err = companyCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	companyStatus = !company.Active
	return
}

// CompanyDelete ...
func CompanyDelete(id string) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctxt       = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Delete company
	_, err = companyCol.DeleteOne(ctxt, filter)
	return
}
