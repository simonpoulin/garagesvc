package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Service ...
type Service struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyID primitive.ObjectID `json:"companyid" bson:"companyid"`
	Name      string             `json:"name" bson:"name"`
	Location  Location           `json:"location" bson:"location"`
	Address   string             `json:"address" bson:"address"`
	Active    bool               `json:"active" bson:"active"`
}

// ServiceCreatePayload ...
type ServiceCreatePayload struct {
	CompanyID       string `json:"companyid" bson:"companyid" valid:"required, stringlength(24|24)"`
	CompanyObjectID primitive.ObjectID
	Name            string   `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Location        Location `json:"location" bson:"location"`
	Address         string   `json:"address" bson:"address" validator:"required, stringlength(1|50)"`
}

// ServiceUpdatePayload ...
type ServiceUpdatePayload struct {
	Name     string   `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Location Location `json:"location" bson:"location"`
	Address  string   `json:"address" bson:"address" validator:"required, stringlength(1|50)"`
	Active   bool     `json:"active" bson:"active" valid:"required, type(bool)"`
}
