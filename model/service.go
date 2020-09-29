package model

import (
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service ...
type Service struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyID   primitive.ObjectID `json:"companyid" bson:"companyid"`
	Name        string             `json:"name" bson:"name"`
	Location    Location           `json:"location" bson:"location"`
	Address     string             `json:"address" bson:"address"`
	Active      bool               `json:"active" bson:"active"`
	Phone       string             `json:"phone" bson:"phone"`
	Email       string             `json:"email" bson:"email"`
	Description string             `json:"desc" bson:"desc"`
}

// ServiceCreateBSON ...
type ServiceCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	CompanyID    primitive.ObjectID `bson:"companyid"`
	Name         string             `bson:"name"`
	Location     Location           `bson:"location"`
	Address      string             `bson:"address"`
	Active       bool               `bson:"active"`
	Phone        string             `bson:"phone"`
	Email        string             `bson:"email"`
	Description  string             `bson:"desc"`
	SearchString string             `bson:"searchstring"`
}

// ServiceUpdateBSON ...
type ServiceUpdateBSON struct {
	Name         string   `bson:"name"`
	Location     Location `bson:"location"`
	Address      string   `bson:"address"`
	Active       bool     `bson:"active"`
	Phone        string   `bson:"phone"`
	Email        string   `bson:"email"`
	Description  string   `bson:"desc"`
	SearchString string   `bson:"searchstring"`
}

// ConvertToCreateBSON ...
func (payload ServiceCreatePayload) ConvertToCreateBSON() (serviceBSON ServiceCreateBSON) {
	serviceBSON = ServiceCreateBSON{
		ID:           primitive.NewObjectID(),
		CompanyID:    payload.CompanyObjectID,
		Name:         payload.Name,
		Location:     payload.Location,
		Address:      payload.Address,
		Active:       false,
		Phone:        payload.Phone,
		Email:        payload.Email,
		Description:  payload.Description,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateBSON ...
func (payload ServiceUpdatePayload) ConvertToUpdateBSON() (serviceBSON ServiceUpdateBSON) {
	serviceBSON = ServiceUpdateBSON{
		Name:         payload.Name,
		Location:     payload.Location,
		Address:      payload.Address,
		Active:       payload.Active,
		Phone:        payload.Phone,
		Email:        payload.Email,
		Description:  payload.Description,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateSearchStringBSON ...
func (s Service) ConvertToUpdateSearchStringBSON() (serviceBSON ServiceUpdateBSON) {
	serviceBSON = ServiceUpdateBSON{
		Name:         s.Name,
		Location:     s.Location,
		Address:      s.Address,
		Active:       s.Active,
		Phone:        s.Phone,
		Email:        s.Email,
		Description:  s.Description,
		SearchString: util.ConvertToHex(s.Name),
	}
	return
}
