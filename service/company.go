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
	var company model.CompanyCreateBSON = payload.ConvertToCreateBSON()

	//Insert to database
	err = dao.CompanyCreate(company)
	companyID = company.ID
	return
}

// CompanyDetail ...
func CompanyDetail(id primitive.ObjectID) (companyRes model.CompanyResponse, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for company from database
	company, err := dao.CompanyFindOne(filter)
	if err != nil {
		return
	}

	companyRes, err = CompanyConvertToResponse(company)

	return
}

// CompanyList ...
func CompanyList(query model.AppQuery) (companyList util.PagedList, err error) {

	var (
		findQuery      = query.GenerateFindQuery()
		companyListRes []model.CompanyResponse
	)

	//Get companies
	companies, err := dao.CompanyFind(findQuery)
	if err != nil {
		return
	}

	for _, company := range companies {
		var companyRes model.CompanyResponse
		companyRes, err = CompanyConvertToResponse(company)
		if err != nil {
			return
		}
		companyListRes = append(companyListRes, companyRes)
	}

	//Paging list
	companyList, err = util.Paging(companyListRes, query.Page, 8)

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

	//Set company filter
	companyFilter := bson.M{"_id": id}

	//Set service query
	query := model.AppQuery{
		CompanyID: id,
	}
	findQuery := query.GenerateFindQuery()

	//Get services
	services, err := dao.ServiceFind(findQuery)
	if err != nil {
		return
	}

	//Delete services
	for _, service := range services {
		err = ServiceDelete(service.ID)
		if err != nil {
			return
		}
	}

	//Delete company
	err = dao.CompanyDelete(companyFilter)

	return
}

// CompanyConvertToResponse ...
func CompanyConvertToResponse(c model.Company) (res model.CompanyResponse, err error) {

	res.ID = c.ID
	res.Name = c.Name
	res.Location = c.Location
	res.Email = c.Email
	res.Address = c.Address
	res.Phone = c.Phone
	res.Active = c.Active

	return
}
