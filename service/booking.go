package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreate ...
func BookingCreate(payload model.BookingCreatePayload, customerID primitive.ObjectID) (bookingID string, err error) {

	var booking model.BookingCreateBSON

	//Set data for new booking
	booking = payload.ConvertToCreateBSON()

	//Insert to database
	err = dao.BookingCreate(booking)
	bookingID = booking.ID.Hex()
	return
}

// BookingDetail ...
func BookingDetail(id primitive.ObjectID) (bookingRes model.BookingResponse, err error) {

	//Set booking filter
	bookingFilter := bson.M{"_id": id}

	//Looking for booking from database
	booking, err := dao.BookingFindOne(bookingFilter)
	if err != nil {
		return
	}

	bookingRes, err = BookingConvertToResponse(booking)

	return
}

// BookingList ...
func BookingList(query model.AppQuery) (bookingList util.PagedList, err error) {
	var (
		findQuery      = query.GenerateFindQuery()
		bookingListRes []model.BookingResponse
	)

	//Get bookings
	bookings, err := dao.BookingFind(findQuery)
	if err != nil {
		return
	}

	//Get booking response list
	for _, booking := range bookings {
		var bookingRes model.BookingResponse
		bookingRes, err = BookingConvertToResponse(booking)
		if err != nil {
			return
		}
		bookingListRes = append(bookingListRes, bookingRes)
	}

	//Paging list
	bookingList, err = util.Paging(bookingListRes, query.Page, 8)

	return
}

// BookingUpdate ...
func BookingUpdate(id primitive.ObjectID, payload model.BookingUpdatePayload) (bookingID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	data := bson.M{"$set": payload.ConvertToUpdateBSON()}

	//Update booking
	err = dao.BookingUpdateOne(filter, data)

	//Return data
	bookingID = id
	return
}

// BookingDelete ...
func BookingDelete(id primitive.ObjectID) (err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Delete booking
	err = dao.BookingDelete(filter)
	return
}

// BookingConvertToResponse ...
func BookingConvertToResponse(b model.Booking) (res model.BookingResponse, err error) {
	res = model.BookingResponse{
		ID:        b.ID,
		Status:    b.Status,
		Date:      b.Date,
		CreatedAt: b.CreatedAt,
		Note:      b.Note,
	}

	//Get service
	service, err := ServiceDetail(b.ServiceID)
	if err != nil {
		return
	}
	res.Service = service

	//Get customer
	customer, err := CustomerDetail(b.CustomerID)
	if err != nil {
		return
	}
	res.Customer = customer

	return
}
