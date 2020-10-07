package model

import (
	"garagesvc/config"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service ...
type Service struct {
	ID          primitive.ObjectID `bson:"_id"`
	CompanyID   primitive.ObjectID `bson:"companyid"`
	Name        string             `bson:"name"`
	Location    Location           `bson:"location"`
	Address     string             `bson:"address"`
	Active      bool               `bson:"active"`
	Phone       string             `bson:"phone"`
	Email       string             `bson:"email"`
	Description string             `bson:"desc"`
	ResourceID  primitive.ObjectID `bson:"resource"`
	SmallImage  string             `bson:"smallimage"`
	MediumImage string             `bson:"mediumimage"`
	LargeImage  string             `bson:"largeimage"`
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
	ResourceID   primitive.ObjectID `bson:"resourceid"`
	SmallImage   string             `bson:"smallimage"`
	MediumImage  string             `bson:"mediumimage"`
	LargeImage   string             `bson:"largeimage"`
}

// ServiceUpdateBSON ...
type ServiceUpdateBSON struct {
	Name         string             `bson:"name"`
	Location     Location           `bson:"location"`
	Address      string             `bson:"address"`
	Active       bool               `bson:"active"`
	Phone        string             `bson:"phone"`
	Email        string             `bson:"email"`
	Description  string             `bson:"desc"`
	SearchString string             `bson:"searchstring"`
	ResourceID   primitive.ObjectID `bson:"resourceid"`
	SmallImage   string             `bson:"smallimage"`
	MediumImage  string             `bson:"mediumimage"`
	LargeImage   string             `bson:"largeimage"`
}

// ConvertToCreateBSON ...
func (payload ServiceCreatePayload) ConvertToCreateBSON() (serviceBSON ServiceCreateBSON) {
	// Get default img name
	img := config.GetIMG()
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
		SmallImage:   img.SmallImage,
		MediumImage:  img.MediumImage,
		LargeImage:   img.LargeImage,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload ServiceUpdatePayload) ConvertToUpdateBSON(resource Resource) (serviceBSON ServiceUpdateBSON) {
	serviceBSON = ServiceUpdateBSON{
		Name:         payload.Name,
		Location:     payload.Location,
		Address:      payload.Address,
		Active:       payload.Active,
		Phone:        payload.Phone,
		Email:        payload.Email,
		Description:  payload.Description,
		SearchString: util.ConvertToHex(payload.Name),
		SmallImage:   resource.SmallImage.GetName(),
		MediumImage:  resource.MediumImage.GetName(),
		LargeImage:   resource.LargeImage.GetName(),
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
