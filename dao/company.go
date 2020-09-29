package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// CompanyCreate ...
func CompanyCreate(company model.CompanyCreateBSON) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	_, err = companyCol.InsertOne(ctx, company)
	return
}

// CompanyFindOne ...
func CompanyFindOne(filter bson.M) (company model.Company, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	err = companyCol.FindOne(ctx, filter).Decode(&company)
	return
}

// CompanyFind ...
func CompanyFind(filter []bson.M) (companyList []model.Company, err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	// Looking for companies
	cur, err := companyCol.Aggregate(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &companyList)
	return
}

// CompanyUpdateOne ...
func CompanyUpdateOne(filter bson.M, data bson.M) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	_, err = companyCol.UpdateOne(ctx, filter, data)
	return
}

// CompanyDelete ...
func CompanyDelete(filter bson.M) (err error) {
	var (
		companyCol = mongodb.CompanyCol()
		ctx        = context.Background()
	)
	_, err = companyCol.DeleteOne(ctx, filter)
	return
}
