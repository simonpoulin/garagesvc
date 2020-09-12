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

// BookingCreatePayload ...
type BookingCreatePayload struct {
	CustomerID       string    `json:"customerid" bson:"customerid" valid:"required, stringlength(24|24)"`
	ServiceID        string    `json:"serviceid" bson:"serviceid" valid:"required, stringlength(24|24)"`
	Date             time.Time `json:"time" bson:"time"`
	Note             string    `json:"note" bson:"note" valid:"stringlength(0|500)"`
	CustomerObjectID primitive.ObjectID
	ServiceObjectID  primitive.ObjectID
}

// BookingUpdatePayload ...
type BookingUpdatePayload struct {
	ServiceID       string    `json:"serviceid" bson:"serviceid" valid:"required, stringlength(24|24)"`
	Date            time.Time `json:"time" bson:"time"`
	Note            string    `json:"note" bson:"note" valid:"stringlength(0|500)"`
	ServiceObjectID primitive.ObjectID
}
