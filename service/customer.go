package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerDetail ...
func CustomerDetail(id primitive.ObjectID) (customerRes model.CustomerResponse, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for customer from database
	customer, err := dao.CustomerFindOne(filter)
	if err != nil {
		return
	}

	customerRes, err = CustomerConvertToResponse(customer)

	return
}

// CustomerList ...
func CustomerList(query model.AppQuery) (customerList util.PagedList, err error) {
	var (
		findQuery      = query.GenerateFindQuery()
		companyListRes []model.CustomerResponse
	)

	//Get customers
	customers, err := dao.CustomerFind(findQuery)
	if err != nil {
		return
	}

	//Get customer response list
	for _, customer := range customers {
		var customerRes model.CustomerResponse
		customerRes, err = CustomerConvertToResponse(customer)
		if err != nil {
			return
		}
		companyListRes = append(companyListRes, customerRes)
	}

	//Paging list
	customerList, err = util.Paging(companyListRes, query.Page, 8)

	return
}

// CustomerUpdate ...
func CustomerUpdate(id primitive.ObjectID, payload model.CustomerUpdatePayload) (customerID primitive.ObjectID, err error) {
	var (
		blankObjectID primitive.ObjectID
	)

	//Set filter
	filter := bson.M{"_id": id}

	//Get resource for image
	rscfilter := bson.M{"_id": payload.ResourceObjectID}
	resource, err := dao.ResourceFindOne(rscfilter)
	if err != nil {
		resource.GetDefaultResource()
	}

	//Get old customer info
	customer, err := dao.CustomerFindOne(filter)

	//Delete old resource
	if customer.ResourceID != blankObjectID && customer.ResourceID != payload.ResourceObjectID {
		ResourceDelete(customer.ResourceID)
	}

	//Check payload password
	if customer.Password != payload.Password {
		payload.Password = util.Hash(payload.Password)
	}

	//Set data
	update := bson.M{"$set": payload.ConvertToUpdateBSON(resource)}

	//Update customer
	err = dao.CustomerUpdateOne(filter, update)

	//Return data
	customerID = id
	return
}

// CustomerDelete ...
func CustomerDelete(id primitive.ObjectID) (err error) {

	//Set customer filter
	customerFilter := bson.M{"_id": id}

	//Set booking query
	query := model.AppQuery{
		CustomerID: id,
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

	//Delete customer
	err = dao.CustomerDelete(customerFilter)
	return
}

// CustomerConvertToResponse ...
func CustomerConvertToResponse(c model.Customer) (res model.CustomerResponse, err error) {
	res = model.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		Phone:       c.Phone,
		Password:    c.Password,
		Address:     c.Address,
		ResourceID:  c.ResourceID,
		SmallImage:  c.SmallImage,
		MediumImage: c.MediumImage,
		LargeImage:  c.LargeImage,
	}

	return
}
