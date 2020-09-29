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
	var service model.ServiceCreateBSON

	//Set data for new service
	service = payload.ConvertToCreateBSON()

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
func ServiceList(query model.AppQuery) (serviceList util.PagedList, err error) {
	var findQuery = query.GenerateFindQuery()

	//Get services
	services, err := dao.ServiceFind(findQuery)
	if err != nil {
		return
	}

	//Paging list
	serviceList, err = util.Paging(services, query.Page, 8)

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
		update = bson.M{"$set": payload.ConvertToUpdateBSON}
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
