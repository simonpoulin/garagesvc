package service

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreate ...
func BookingCreate(payload model.BookingCreatePayload, customerID interface{}) (bookingID string, err error) {
	var booking model.Booking

	//Set data for new booking
	if customerID != nil {
		booking.CustomerID = customerID.(primitive.ObjectID)
	} else {
		booking.CustomerID = payload.CustomerObjectID
	}
	booking.ServiceID = payload.ServiceObjectID
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
func BookingList(status string, serviceID primitive.ObjectID, customerID primitive.ObjectID, page int) (bookingList interface{}, err error) {
	var (
		filterParts []bson.M
		findQuery   []bson.M
	)

	//Set filter parts
	if status != "" {
		stt, _ := strconv.ParseBool(status)
		filterParts = append(filterParts, bson.M{"status": stt})
	}

	if serviceID.Hex() != "000000000000000000000000" {
		filterParts = append(filterParts, bson.M{"serviceid": serviceID})
	}

	if customerID.Hex() != "000000000000000000000000" {
		filterParts = append(filterParts, bson.M{"customerid": customerID})
	}

	//Set filter query from parts
	findQuery = append(findQuery, bson.M{"$match": func() bson.M {
		if filterParts != nil {
			if len(filterParts) > 0 {
				return bson.M{"$and": filterParts}
			}
		}
		return bson.M{}
	}()})

	//Get bookings
	bookings, err := dao.BookingFind(findQuery)

	//Paging list
	if page > 0 {
		bookingList, err = util.Paging(bookings, page, 8)
		return
	}
	bookingList = bookings

	return
}

// BookingUpdate ...
func BookingUpdate(id primitive.ObjectID, payload model.BookingUpdatePayload) (bookingID primitive.ObjectID, err error) {

	//Set filter and data
	filter := bson.M{"_id": id}
	data := bson.M{"$set": bson.M{
		"serviceid": payload.ServiceObjectID,
		"date":      payload.Date,
		"note":      payload.Note,
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
