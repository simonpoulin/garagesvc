package service

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreate ...
func BookingCreate(payload model.BookingCreatePayload) (bookingID string, err error) {
	var (
		booking    model.Booking
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set data for new booking
	booking.CustomerID, err = primitive.ObjectIDFromHex(payload.CustomerID)
	if err != nil {
		return
	}
	booking.ServiceID, err = primitive.ObjectIDFromHex(payload.ServiceID)
	if err != nil {
		return
	}
	booking.ID = primitive.NewObjectID()
	booking.Status = "Pending"
	booking.Date = payload.Date
	booking.Note = payload.Note

	//Insert to database
	_, err = bookingCol.InsertOne(ctxt, booking)
	bookingID = booking.ID.Hex()
	return
}

// BookingDetail ...
func BookingDetail(id string) (e model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Looking for booking from database
	err = bookingCol.FindOne(ctxt, filter).Decode(&e)
	return
}

// BookingList ...
func BookingList() (bookingList []model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Get bookings
	cur, err := bookingCol.Find(ctxt, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add bookings to list
	for cur.Next(ctxt) {
		var result model.Booking
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		bookingList = append(bookingList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// BookingListByStatus ...
func BookingListByStatus(status string) (bookingList []model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter
	filter := bson.M{"status": status}

	//Get bookings
	cur, err := bookingCol.Find(ctxt, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add bookings to list
	for cur.Next(ctxt) {
		var result model.Booking
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		bookingList = append(bookingList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// BookingListByServiceID ...
func BookingListByServiceID(serviceID string) (bookingList []model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter
	svcID, err := primitive.ObjectIDFromHex(serviceID)
	if err != nil {
		return
	}
	filter := bson.M{"service_id": svcID}

	//Get bookings
	cur, err := bookingCol.Find(ctxt, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add bookings to list
	for cur.Next(ctxt) {
		var result model.Booking
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		bookingList = append(bookingList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// BookingListByCustomerID ...
func BookingListByCustomerID(customerID string) (bookingList []model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter
	ctmID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return
	}
	filter := bson.M{"customer_id": ctmID}

	//Get bookings
	cur, err := bookingCol.Find(ctxt, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctxt)

	//Add bookings to list
	for cur.Next(ctxt) {
		var result model.Booking
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		bookingList = append(bookingList, result)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return
}

// BookingUpdate ...
func BookingUpdate(id string, payload model.BookingUpdatePayload) (bookingID string, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter and data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"service_id": payload.ServiceID,
		"date":       payload.Date,
		"note":       payload.Note,
	}}

	//Update booking
	_, err = bookingCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	bookingID = id
	return
}

// BookingChangeStatus ...
func BookingChangeStatus(id string, status string) (bookingStatus string, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter and active state data
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"status": status}}

	//Update booking
	_, err = bookingCol.UpdateOne(ctxt, filter, update)
	if err != nil {
		return
	}

	//Return data
	bookingStatus = status
	return
}

// BookingDelete ...
func BookingDelete(id string) (err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctxt       = context.Background()
	)

	//Set filter
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": _id}

	//Delete booking
	_, err = bookingCol.DeleteOne(ctxt, filter)
	return
}
