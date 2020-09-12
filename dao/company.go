package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"
)

// CompanyCreate ...
func CompanyCreate(company model.Company) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	_, err = companyCol.InsertOne(ctx, company)
	return
}

// CompanyFindOne ...
func CompanyFindOne(filter interface{}) (company model.Company, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	err = companyCol.FindOne(ctx, filter).Decode(&company)
	return
}

// CompanyFind ...
func CompanyFind(filter interface{}) (companyList []model.Company, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	// Looking for companys
	cur, err := companyCol.Find(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &companyList)
	return
}

// CompanyUpdateOne ...
func CompanyUpdateOne(filter interface{}, data interface{}) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	_, err = companyCol.UpdateOne(ctx, filter, data)
	return
}

// CompanyDelete ...
func CompanyDelete(filter interface{}) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	_, err = companyCol.DeleteOne(ctx, filter)
	return
}
