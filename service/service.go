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
func ServiceDetail(id primitive.ObjectID) (serviceRes model.ServiceResponse, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for service from database
	service, err := dao.ServiceFindOne(filter)
	if err != nil {
		return
	}

	serviceRes, err = ServiceConvertToResponse(service)

	return
}

// ServiceList ...
func ServiceList(query model.AppQuery) (serviceList util.PagedList, err error) {
	var (
		findQuery      = query.GenerateFindQuery()
		serviceListRes []model.ServiceResponse
	)

	//Get services
	services, err := dao.ServiceFind(findQuery)
	if err != nil {
		return
	}

	//Get service response list
	for _, service := range services {
		var serviceRes model.ServiceResponse
		serviceRes, err = ServiceConvertToResponse(service)
		if err != nil {
			return
		}
		serviceListRes = append(serviceListRes, serviceRes)
	}

	//Paging list
	serviceList, err = util.Paging(serviceListRes, query.Page, 8)

	return
}

// ServiceUpdate ...
func ServiceUpdate(id primitive.ObjectID, payload model.ServiceUpdatePayload, active string) (serviceID primitive.ObjectID, err error) {

	var (
		update        bson.M
		blankObjectID primitive.ObjectID
	)

	//Set filter and data
	filter := bson.M{"_id": id}

	//Get resource for image
	rscfilter := bson.M{"_id": payload.ResourceObjectID}
	resource, err := dao.ResourceFindOne(rscfilter)
	if err != nil {
		resource.GetDefaultResource()
	}

	//Get old service info
	service, err := dao.ServiceFindOne(filter)

	//Delete old resource
	if service.ResourceID != blankObjectID && service.ResourceID != payload.ResourceObjectID {
		ResourceDelete(service.ResourceID)
	}

	if active != "" {
		stt, _ := strconv.ParseBool(active)
		update = bson.M{"$set": bson.M{"active": stt}}
	} else {
		update = bson.M{"$set": payload.ConvertToUpdateBSON(resource)}
	}

	//Update service
	err = dao.ServiceUpdateOne(filter, update)

	//Return data
	serviceID = id
	return
}

// ServiceDelete ...
func ServiceDelete(id primitive.ObjectID) (err error) {

	//Set service filter
	serviceFilter := bson.M{"_id": id}

	//Set booking query
	query := model.AppQuery{
		ServiceID: id,
	}
	findQuery := query.GenerateFindQuery()

	//Get bookings
	bookings, err := dao.BookingFind(findQuery)
	if err != nil {
		return
	}

	//Delete bookings
	for _, booking := range bookings {
		err = BookingDelete(booking.ID)
		if err != nil {
			return
		}
	}

	//Delete service
	err = dao.ServiceDelete(serviceFilter)
	return
}

// ServiceConvertToResponse ...
func ServiceConvertToResponse(s model.Service) (res model.ServiceResponse, err error) {
	res = model.ServiceResponse{
		ID:          s.ID,
		Name:        s.Name,
		Location:    s.Location,
		Address:     s.Address,
		Active:      s.Active,
		Phone:       s.Phone,
		Email:       s.Email,
		Description: s.Description,
		ResourceID:  s.ResourceID,
		SmallImage:  s.SmallImage,
		MediumImage: s.MediumImage,
		LargeImage:  s.LargeImage,
	}

	//Get company
	company, err := CompanyDetail(s.CompanyID)
	if err != nil {
		return
	}
	res.Company = company

	return
}
