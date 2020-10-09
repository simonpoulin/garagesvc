package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// ServiceCreatePayload ...
type ServiceCreatePayload struct {
	CompanyID       string   `json:"companyid" valid:"required, stringlength(24|24)"`
	Name            string   `json:"name" valid:"required, stringlength(1|50)"`
	Location        Location `json:"location"`
	Address         string   `json:"address" valid:"required, stringlength(1|200)"`
	Phone           string   `json:"phone" valid:"type(string), stringlength(10|10)"`
	Email           string   `json:"email"`
	Description     string   `json:"description" valid:"stringlength(0|500)"`
	CompanyObjectID primitive.ObjectID
}

// ServiceUpdatePayload ...
type ServiceUpdatePayload struct {
	Name             string   `json:"name" valid:"required, stringlength(1|50)"`
	Location         Location `json:"location" `
	Address          string   `json:"address" valid:"required, stringlength(1|200)"`
	Active           bool     `json:"active" valid:"required, type(bool)"`
	Phone            string   `json:"phone" valid:"type(string), stringlength(10|10)"`
	Email            string   `json:"email"`
	Description      string   `json:"description" valid:"stringlength(0|500)"`
	ResourceID       string   `json:"resourceid"  valid:"stringlength(24|24)"`
	ResourceObjectID primitive.ObjectID
}

// ServiceQuery ...
type ServiceQuery struct {
	Name            string `query:"name"`
	CompanyID       string `query:"companyid"`
	Active          string `query:"active"`
	Page            int    `query:"page"`
	Phone           string `query:"page"`
	CompanyObjectID primitive.ObjectID
}
