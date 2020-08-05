package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyID primitive.ObjectID `json:"company_id" bson:"company_id" validator:"eq=24"`
	Name      string             `json:"name" bson:"name" validator:"required, max=20, alphaunicode"`
	Location  string             `json:"location" bson:"location"`
	Address   string             `json:"address" bson:"address" validator:"required, max=50"`
	Active    bool               `json:"active" bson:"active"`
}
