package service

import (
	"context"
	"errors"
	"garagesvc/model"
	"garagesvc/module/mongodb"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerCreate ...
func CustomerCreate(payload model.CustomerPayload) (customerID string, err error) {
	var (
		customer    model.Customer
		customerCol = mongodb.CustomerCol()
		ctxt        = context.Background()
	)

	//Set data for new customer
	customer.ID = primitive.NewObjectID()
	customer.Password = util.Hash(payload.Password)
	customer.Name = payload.Name
	customer.Phone = payload.Phone

	//Insert to database
	_, err = customerCol.InsertOne(ctxt, customer)
	customerID = customer.ID.Hex()
	return
}

// CustomerLogin ...
func CustomerLogin(payload model.CustomerLoginPayload) (token string, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctxt        = context.Background()
	)

	//Get customer by phone number
	payload.Password = util.Hash(payload.Password)
	var e model.Customer
	filter := bson.M{"phone": payload.Phone}
	err = customerCol.FindOne(ctxt, filter).Decode(&e)
	if err != nil {
		return
	}

	//Check password match
	if payload.Password != e.Password {
		err = errors.New("password not match")
		return
	}
	token, err = util.TokenEncode(e.ID.Hex())
	return
}

// CustomerDetail ...
func CustomerDetail(id string) (e model.Customer, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctxt        = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Looking for customer from database
	err = customerCol.FindOne(ctxt, filter).Decode(&e)
	return
}

// CustomerList ...
func CustomerList() (customerList []model.Customer, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctxt        = context.Background()
	)

	//Get customers
	cur, err := customerCol.Find(ctxt, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add customers to list
	for cur.Next(ctxt) {
		var result model.Customer
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		customerList = append(customerList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// CustomerUpdate ...
func CustomerUpdate(id string, payload model.CustomerPayload) (customerID string, err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctxt        = context.Background()
	)

	//Set filter and data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"password": payload.Password,
		"name":     payload.Name,
		"phone":    payload.Phone,
	}}

	//Update customer
	_, err = customerCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	customerID = id
	return
}

// CustomerDelete ...
func CustomerDelete(id string) (err error) {
	var (
		customerCol = mongodb.CustomerCol()
		ctxt        = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Delete customer
	_, err = customerCol.DeleteOne(ctxt, filter)
	return
}
