package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Company ...
type Company struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Location Location           `json:"location" bson:"location"`
	Address  string             `json:"address" bson:"address"`
	Active   bool               `json:"active" bson:"active"`
}

// CompanyCreatePayload ...
type CompanyCreatePayload struct {
	Name     string   `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Location Location `json:"location" bson:"location"`
	Address  string   `json:"address" bson:"address" validator:"required, stringlength(1|50)"`
}

// CompanyUpdatePayload ...
type CompanyUpdatePayload struct {
	Name     string   `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Location Location `json:"location" bson:"location"`
	Address  string   `json:"address" bson:"address" validator:"required, stringlength(1|50)"`
	Active   bool     `json:"active" bson:"active" valid:"required, type(bool)"`
}
