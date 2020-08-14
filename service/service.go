package service

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceCreate ...
func ServiceCreate(companyID string, payload model.ServiceCreatePayload) (serviceID string, err error) {
	var (
		service    model.Service
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Set data for new service
	service.CompanyID, err = primitive.ObjectIDFromHex(companyID)
	if err != nil {
		return
	}
	service.ID = primitive.NewObjectID()
	service.Location = payload.Location
	service.Active = false
	service.Address = payload.Address
	service.Name = payload.Name

	//Insert to database
	_, err = serviceCol.InsertOne(ctxt, service)
	serviceID = service.ID.Hex()
	return
}

// ServiceDetail ...
func ServiceDetail(id string) (e model.Service, err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Looking for service from database
	err = serviceCol.FindOne(ctxt, filter).Decode(&e)
	return
}

// ServiceList ...
func ServiceList(companyID string) (serviceList []model.Service, err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Get services by company ID
	filter := bson.M{"company_id": companyID}
	cur, err := serviceCol.Find(ctxt, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add services to list
	for cur.Next(ctxt) {
		var result model.Service
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		serviceList = append(serviceList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// ServiceListByActiveState ...
func ServiceListByActiveState(active string) (serviceList []model.Service, err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Set filter
	filter := bson.M{"active": active}

	//Get services

	cur, err := serviceCol.Find(ctxt, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add services to list
	for cur.Next(ctxt) {
		var result model.Service
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		serviceList = append(serviceList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// ServiceUpdate ...
func ServiceUpdate(id string, payload model.ServiceUpdatePayload) (serviceID string, err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Set filter and data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"active":   payload.Active,
		"address":  payload.Address,
		"name":     payload.Name,
		"location": payload.Location,
	}}

	//Update service
	_, err = serviceCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	serviceID = id
	return
}

// ServiceChangeActive ...
func ServiceChangeActive(id string) (serviceStatus bool, err error) {
	var (
		service    model.Service
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Set filter and data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"active": !service.Active}}

	//Update service
	_, err = serviceCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	serviceStatus = !service.Active
	return
}

// ServiceDelete ...
func ServiceDelete(id string) (err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctxt       = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Delete service
	_, err = serviceCol.DeleteOne(ctxt, filter)
	return
}
