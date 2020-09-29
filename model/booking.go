package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Booking ...
type Booking struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	CustomerID primitive.ObjectID `json:"customerid" bson:"customerid"`
	ServiceID  primitive.ObjectID `json:"serviceid" bson:"serviceid"`
	Status     string             `json:"status" bson:"status"`
	Date       time.Time          `json:"time" bson:"time"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	Note       string             `json:"note" bson:"note"`
}

// BookingCreateBSON ...
type BookingCreateBSON struct {
	ID         primitive.ObjectID `bson:"_id"`
	CustomerID primitive.ObjectID `bson:"customerid"`
	ServiceID  primitive.ObjectID `bson:"serviceid"`
	Status     string             `bson:"status"`
	Date       time.Time          `bson:"time"`
	CreatedAt  time.Time          `bson:"createdAt"`
	Note       string             `bson:"note"`
}

// BookingUpdateBSON ...
type BookingUpdateBSON struct {
	ServiceID primitive.ObjectID `bson:"serviceid"`
	Date      time.Time          `bson:"time"`
	Note      string             `bson:"note"`
}

// ConvertToCreateBSON ...
func (payload BookingCreatePayload) ConvertToCreateBSON() (bookingBSON BookingCreateBSON) {
	bookingBSON = BookingCreateBSON{
		ID:         primitive.NewObjectID(),
		CustomerID: payload.CustomerObjectID,
		ServiceID:  payload.ServiceObjectID,
		Status:     "Pending",
		Date:       payload.Date,
		CreatedAt:  time.Now(),
		Note:       payload.Note,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload BookingUpdatePayload) ConvertToUpdateBSON() (bookingBSON BookingUpdateBSON) {
	bookingBSON = BookingUpdateBSON{
		ServiceID: payload.ServiceObjectID,
		Date:      payload.Date,
		Note:      payload.Note,
	}
	return
}
