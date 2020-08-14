package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Booking ...
type Booking struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ServiceID primitive.ObjectID `json:"service_id" bson:"service_id"`
	Status    string             `json:"status" bson:"status"`
	Date      time.Time          `json:"time" bson:"time"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Note      string             `json:"note" bson:"note"`
}

// BookingCreatePayload ...
type BookingCreatePayload struct {
	UserID    string    `json:"user_id" bson:"user_id" valid:"required, stringlength(24|24)"`
	ServiceID string    `json:"service_id" bson:"service_id" valid:"required, stringlength(24|24)"`
	Date      time.Time `json:"time" bson:"time"`
	Note      string    `json:"note" bson:"note" valid:"stringlength(0|500)"`
}

// BookingUpdatePayload ...
type BookingUpdatePayload struct {
	ServiceID string    `json:"service_id" bson:"service_id" valid:"required, stringlength(24|24)"`
	Date      time.Time `json:"time" bson:"time"`
	Note      string    `json:"note" bson:"note" valid:"stringlength(0|500)"`
}
