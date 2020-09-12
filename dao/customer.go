package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"
)

// CustomerCreate ...
func CustomerCreate(customer model.Customer) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	_, err = customerCol.InsertOne(ctx, customer)
	return
}

// CustomerFindOne ...
func CustomerFindOne(filter interface{}) (customer model.Customer, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	err = customerCol.FindOne(ctx, filter).Decode(&customer)
	return
}

// CustomerFind ...
func CustomerFind(filter interface{}) (customerList []model.Customer, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	// Looking for customers
	cur, err := customerCol.Find(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &customerList)
	return
}

// CustomerUpdateOne ...
func CustomerUpdateOne(filter interface{}, data interface{}) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	_, err = customerCol.UpdateOne(ctx, filter, data)
	return
}

// CustomerDelete ...
func CustomerDelete(filter interface{}) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctx         = context.Background()
	)
	_, err = customerCol.DeleteOne(ctx, filter)
	return
}
