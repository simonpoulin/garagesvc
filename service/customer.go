package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerDetail ...
func CustomerDetail(id primitive.ObjectID) (customer model.Customer, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for customer from database
	customer, err = dao.CustomerFindOne(filter)
	return
}

// CustomerList ...
func CustomerList(name string, page int) (customerList util.PagedList, err error) {
	var (
		filter = bson.M{}
	)

	if name != "" {
		filter = bson.M{"name": bson.M{"$regex": name}}
	}

	//Get customers
	customers, err := dao.CustomerFind(filter)
	if err != nil {
		return
	}

	//Paging list
	customerList, err = util.Paging(customers, page, 8)

	return
}

// CustomerUpdate ...
func CustomerUpdate(id primitive.ObjectID, payload model.CustomerUpdatePayload) (customerID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"password": util.Hash(payload.Password),
		"name":     payload.Name,
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
