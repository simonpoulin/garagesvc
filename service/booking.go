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
func BookingDetail(id primitive.ObjectID) (booking model.Booking, err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Looking for booking from database
	booking, err = dao.BookingFindOne(filter)
	return
}

// BookingList ...
func BookingList(query model.AppQuery) (bookingList util.PagedList, err error) {
	var findQuery = query.GenerateFindQuery()

	//Get bookings
	bookings, err := dao.BookingFind(findQuery)
	if err != nil {
		return
	}

	//Paging list
	bookingList, err = util.Paging(bookings, query.Page, 8)

	return
}

// BookingUpdate ...
func BookingUpdate(id primitive.ObjectID, payload model.BookingUpdatePayload) (bookingID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	data := bson.M{"$set": payload.ConvertToUpdateBSON}

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
