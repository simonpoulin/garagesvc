package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceCreate ...
func ServiceCreate(payload model.ServiceCreatePayload) (serviceID primitive.ObjectID, err error) {
	var service model.Service

	//Set data for new service
	service.CompanyID = payload.CompanyObjectID
	service.ID = primitive.NewObjectID()
	service.Location = payload.Location
	service.Active = false
	service.Address = payload.Address
	service.Name = payload.Name

	//Insert to database
	err = dao.ServiceCreate(service)
	serviceID = service.ID
	return
}

// ServiceDetail ...
func ServiceDetail(id primitive.ObjectID) (service model.Service, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for service from database
	service, err = dao.ServiceFindOne(filter)
	return
}

// ServiceList ...
func ServiceList(active string, name string, companyID primitive.ObjectID, page int) (serviceList interface{}, err error) {
	var (
		filterParts []bson.M
		findQuery   []bson.M
	)

	//Filter parts
	if active != "" {
		filterParts = append(filterParts, bson.M{"active": active})
	}

	if name != "" {
		filterParts = append(filterParts, bson.M{"name": bson.M{"$regex": name}})
	}

	if companyID.Hex() != "000000000000000000000000" {
		filterParts = append(filterParts, bson.M{"companyid": companyID})
	}

	//Getting filter query
	findQuery = append(findQuery, bson.M{"$match": func() bson.M {
		if filterParts != nil {
			if len(filterParts) > 0 {
				return bson.M{"$and": filterParts}
			}
		}
		return bson.M{}
	}()})

	//Get services
	services, err := dao.ServiceFind(findQuery)

	//Paging list
	if page > 0 {
		serviceList, err = util.Paging(services, page, 8)
		return
	}
	serviceList = services

	return
}

// ServiceUpdate ...
func ServiceUpdate(id primitive.ObjectID, payload model.ServiceUpdatePayload, active string) (serviceID primitive.ObjectID, err error) {

	var update bson.M

	//Set filter and data
	filter := bson.M{"_id": id}

	if active != "" {
		stt, _ := strconv.ParseBool(active)
		update = bson.M{"$set": bson.M{"active": stt}}
	} else {
		update = bson.M{"$set": bson.M{
			"active":   payload.Active,
			"address":  payload.Address,
			"name":     payload.Name,
			"location": payload.Location,
		}}
	}

	//Update service
	err = dao.ServiceUpdateOne(filter, update)

	//Return data
	serviceID = id
	return
}

// ServiceDelete ...
func ServiceDelete(id primitive.ObjectID) (err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Delete service
	err = dao.ServiceDelete(filter)
	return
}
