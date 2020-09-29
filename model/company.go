package model

import (
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Company ...
type Company struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Location Location           `json:"location" bson:"location"`
	Email    string             `json:"email" bson:"email"`
	Address  string             `json:"address" bson:"address"`
	Phone    string             `json:"phone" bson:"phone"`
	Active   bool               `json:"active" bson:"active"`
}

//CompanyCreateBSON ...
type CompanyCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Location     Location           `bson:"location"`
	Email        string             `bson:"email"`
	Address      string             `bson:"address"`
	Active       bool               `bson:"active"`
	Phone        string             `bson:"phone"`
	SearchString string             `bson:"searchstring"`
}

//CompanyUpdateBSON ...
type CompanyUpdateBSON struct {
	Name         string   `bson:"name"`
	Location     Location `bson:"location"`
	Email        string   `bson:"email"`
	Address      string   `bson:"address"`
	Active       bool     `bson:"active"`
	Phone        string   `bson:"phone"`
	SearchString string   `bson:"searchstring"`
}

// ConvertToCreateBSON ...
func (payload CompanyCreatePayload) ConvertToCreateBSON() (companyBSON CompanyCreateBSON) {
	companyBSON = CompanyCreateBSON{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Location:     payload.Location,
		Email:        payload.Email,
		Address:      payload.Address,
		Active:       false,
		Phone:        payload.Phone,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateBSON ...
func (payload CompanyUpdatePayload) ConvertToUpdateBSON() (companyBSON CompanyUpdateBSON) {
	companyBSON = CompanyUpdateBSON{
		Name:         payload.Name,
		Location:     payload.Location,
		Email:        payload.Email,
		Address:      payload.Address,
		Active:       payload.Active,
		Phone:        payload.Phone,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateSearchStringBSON ...
func (c Company) ConvertToUpdateSearchStringBSON() (companyBSON CompanyUpdateBSON) {
	companyBSON = CompanyUpdateBSON{
		Name:         c.Name,
		Location:     c.Location,
		Address:      c.Address,
		Active:       c.Active,
		Phone:        c.Phone,
		SearchString: util.ConvertToHex(c.Name),
	}
	return
}
