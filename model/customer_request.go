package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CustomerLoginPayload ...
type CustomerLoginPayload struct {
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
}

// CustomerRegisterPayload ...
type CustomerRegisterPayload struct {
	Name     string `json:"name" valid:"required, stringlength(1|50)"`
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
	Address  string `json:"address" valid:"stringlength(1|200)"`
}

// CustomerUpdatePayload ...
type CustomerUpdatePayload struct {
	Name             string `json:"name" valid:"required, stringlength(1|50)"`
	Password         string `json:"password" valid:"required, type(string), stringlength(6|50)"`
	Address          string `json:"address" valid:"stringlength(1|200)"`
	ResourceID       string `json:"resourceid"  valid:"stringlength(24|24)"`
	ResourceObjectID primitive.ObjectID
}

// CustomerQuery ...
type CustomerQuery struct {
	Name string `query:"name"`
	Page int    `query:"page"`
}
