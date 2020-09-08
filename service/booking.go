package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreate ...
func BookingCreate(payload model.BookingCreatePayload) (bookingID string, err error) {
	var booking model.Booking

	//Set data for new booking
	booking.CustomerID = payload.CustomerID
	booking.ServiceID = payload.ServiceID
	booking.ID = primitive.NewObjectID()
	booking.Status = "Pending"
	booking.Date = payload.Date
	booking.Note = payload.Note
	booking.CreatedAt = time.Now()

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
func BookingList() (bookingList []model.Booking, err error) {

	//Get bookings
	bookingList, err = dao.BookingFind(bson.M{})
	return
}

// BookingListByStatus ...
func BookingListByStatus(status string) (bookingList []model.Booking, err error) {

	//Set filter
	filter := bson.M{"status": status}

	//Get bookings
	bookingList, err = dao.BookingFind(filter)
	return
}

// BookingListByServiceID ...
func BookingListByServiceID(serviceID primitive.ObjectID) (bookingList []model.Booking, err error) {

	//Set filter
	filter := bson.M{"service_id": serviceID}

	//Get bookings
	bookingList, err = dao.BookingFind(filter)
	return
}

// BookingListByCustomerID ...
func BookingListByCustomerID(customerID primitive.ObjectID) (bookingList []model.Booking, err error) {

	//Set filter
	filter := bson.M{"customer_id": customerID}

	//Get bookings
	bookingList, err = dao.BookingFind(filter)
	return
}

// BookingUpdate ...
func BookingUpdate(id primitive.ObjectID, payload model.BookingUpdatePayload) (bookingID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	data := bson.M{"$set": bson.M{
		"service_id": payload.ServiceID,
		"date":       payload.Date,
		"note":       payload.Note,
	}}

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
