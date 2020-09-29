package service

import (
	"errors"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerRegister ...
func CustomerRegister(payload model.CustomerRegisterPayload) (customerID primitive.ObjectID, err error) {
	var customer model.CustomerCreateBSON

	//Set data for new customer
	customer = payload.ConvertToCreateBSON()

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
