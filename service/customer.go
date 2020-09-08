package service

import (
	"errors"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerCreate ...
func CustomerCreate(payload model.CustomerPayload) (customerID primitive.ObjectID, err error) {
	var customer model.Customer

	//Set data for new customer
	customer.ID = primitive.NewObjectID()
	customer.Password = util.Hash(payload.Password)
	customer.Name = payload.Name
	customer.Phone = payload.Phone

	//Insert to database
	err = dao.CustomerCreate(customer)
	customerID = customer.ID
	return
}

// CustomerLogin ...
func CustomerLogin(payload model.CustomerLoginPayload) (token string, err error) {

	//Get customer by phone number
	payload.Password = util.Hash(payload.Password)
	filter := bson.M{"phone": payload.Phone}
	customer, err := dao.CustomerFindOne(filter)
	if err != nil {
		return
	}

	//Check password match
	if payload.Password != customer.Password {
		err = errors.New("password not match")
		return
	}
	token, err = customer.GenerateToken()
	return
}

// CustomerDetail ...
func CustomerDetail(id primitive.ObjectID) (customer model.Customer, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for customer from database
	customer, err = dao.CustomerFindOne(filter)
	return
}

// CustomerList ...
func CustomerList() (customerList []model.Customer, err error) {

	//Get customers
	customerList, err = dao.CustomerFind(bson.M{})
	return
}

// CustomerUpdate ...
func CustomerUpdate(id primitive.ObjectID, payload model.CustomerPayload) (customerID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"password": payload.Password,
		"name":     payload.Name,
		"phone":    payload.Phone,
	}}

	//Update customer
	err = dao.CustomerUpdateOne(filter, update)

	//Return data
	customerID = id
	return
}

// CustomerDelete ...
func CustomerDelete(id primitive.ObjectID) (err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Delete customer
	err = dao.CustomerDelete(filter)
	return
}
