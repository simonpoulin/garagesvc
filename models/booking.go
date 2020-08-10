package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id" validator:"eq=24"`
	ServiceID primitive.ObjectID `json:"service_id" bson:"service_id" validator:"eq=24"`
	StatusID  int                `json:"status_id" bson:"status_id"`
	Date      time.Time          `json:"time" bson:"time"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Note      string             `json:"note" bson:"note" validator:"max=500"`
}
