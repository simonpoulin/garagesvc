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
	var company model.CompanyCreateBSON

	//Set data for new company
	company = payload.ConvertToCreateBSON()

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
func CompanyList(query model.AppQuery) (companyList util.PagedList, err error) {

	var findQuery = query.GenerateFindQuery()

	//Get companies
	companies, err := dao.CompanyFind(findQuery)
	if err != nil {
		return
	}

	//Paging list
	companyList, err = util.Paging(companies, query.Page, 8)

	return
}

// CompanyUpdate ...
func CompanyUpdate(id primitive.ObjectID, payload model.CompanyUpdatePayload, active string) (companyID primitive.ObjectID, err error) {

	var (
		companyData model.CompanyUpdateBSON
		update      bson.M
	)

	//Set filter and data
	filter := bson.M{"_id": id}

	if active != "" {
		stt, _ := strconv.ParseBool(active)
		update = bson.M{"$set": bson.M{"active": stt}}
	} else {
		companyData = payload.ConvertToUpdateBSON()
		update = bson.M{"$set": companyData}
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
