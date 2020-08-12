package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Company ...
type Company struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name" validator:"required, max=20, alphaunicode"`
	Location Location           `json:"location" bson:"location"`
	Address  string             `json:"address" bson:"address" validator:"required, max=50"`
	Active   bool               `json:"active" bson:"active"`
}
