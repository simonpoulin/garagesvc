package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// CustomerCreate ...
func CustomerCreate(customer model.CustomerCreateBSON) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	_, err = customerCol.InsertOne(ctx, customer)
	return
}

// CustomerFindOne ...
func CustomerFindOne(filter bson.M) (customer model.Customer, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	err = customerCol.FindOne(ctx, filter).Decode(&customer)
	return
}

// CustomerFind ...
func CustomerFind(filter []bson.M) (customerList []model.Customer, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	// Looking for customers
	cur, err := customerCol.Aggregate(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &customerList)
	return
}

// CustomerUpdateOne ...
func CustomerUpdateOne(filter bson.M, data bson.M) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	_, err = customerCol.UpdateOne(ctx, filter, data)
	return
}

// CustomerDelete ...
func CustomerDelete(filter bson.M) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	_, err = customerCol.DeleteOne(ctx, filter)
	return
}
