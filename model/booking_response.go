package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingResponse ...
type BookingResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Customer  CustomerResponse   `json:"customer"`
	Service   ServiceResponse    `json:"service"`
	Status    string             `json:"status"`
	Date      time.Time          `json:"time"`
	CreatedAt time.Time          `json:"createdAt"`
	Note      string             `json:"note"`
}
