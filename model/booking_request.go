package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreatePayload ...
type BookingCreatePayload struct {
	CustomerID       string    `json:"customerid" valid:"required, stringlength(24|24)"`
	ServiceID        string    `json:"serviceid" valid:"required, stringlength(24|24)"`
	Date             time.Time `json:"time"`
	Note             string    `json:"note" valid:"stringlength(0|500)"`
	CustomerObjectID primitive.ObjectID
	ServiceObjectID  primitive.ObjectID
}

// BookingUpdatePayload ...
type BookingUpdatePayload struct {
	ServiceID       string    `json:"serviceid" valid:"required, stringlength(24|24)"`
	Date            time.Time `json:"time"`
	Note            string    `json:"note" valid:"stringlength(0|500)"`
	ServiceObjectID primitive.ObjectID
}

// BookingQuery ...
type BookingQuery struct {
	Status           string `query:"status"`
	ServiceID        string `query:"serviceid"`
	CustomerID       string `query:"customerid"`
	Page             int    `query:"page"`
	ServiceObjectID  primitive.ObjectID
	CustomerObjectID primitive.ObjectID
}
