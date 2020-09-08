package service

import (
	"garagesvc/dao"
	"garagesvc/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceCreate ...
func ServiceCreate(payload model.ServiceCreatePayload) (serviceID primitive.ObjectID, err error) {
	var service model.Service

	//Set data for new service
	service.CompanyID = payload.CompanyID
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
func ServiceList() (serviceList []model.Service, err error) {

	//Get services by company ID
	serviceList, err = dao.ServiceFind(bson.M{})
	return
}

// ServiceListByCompanyID ...
func ServiceListByCompanyID(companyID primitive.ObjectID) (serviceList []model.Service, err error) {

	//Set filter
	filter := bson.M{"company_id": companyID}

	//Get services by company ID
	serviceList, err = dao.ServiceFind(filter)
	return
}

// ServiceListByActiveState ...
func ServiceListByActiveState(active string) (serviceList []model.Service, err error) {

	//Set filter
	filter := bson.M{"active": active}

	//Get services
	serviceList, err = dao.ServiceFind(filter)
	return
}

// ServiceUpdate ...
func ServiceUpdate(id primitive.ObjectID, payload model.ServiceUpdatePayload) (serviceID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"active":   payload.Active,
		"address":  payload.Address,
		"name":     payload.Name,
		"location": payload.Location,
	}}

	//Update service
	err = dao.ServiceUpdateOne(filter, update)

	//Return data
	serviceID = id
	return
}

// ServiceChangeActive ...
func ServiceChangeActive(id primitive.ObjectID) (serviceID primitive.ObjectID, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Get service
	service, err := dao.ServiceFindOne(filter)
	if err != nil {
		return
	}

	//Set active state data
	update := bson.M{"$set": bson.M{"active": !service.Active}}

	//Update service
	err = dao.ServiceUpdateOne(filter, update)

	//Return data
	serviceID = service.ID
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
