package model

import (
	"garagesvc/config"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer ...
type Customer struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Phone       string             `bson:"phone"`
	Password    string             `bson:"password"`
	Address     string             `bson:"address"`
	ResourceID  primitive.ObjectID `bson:"resourceid"`
	SmallImage  string             `bson:"smallimage"`
	MediumImage string             `bson:"mediumimage"`
	LargeImage  string             `bson:"largeimage"`
}

// CustomerCreateBSON ...
type CustomerCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Phone        string             `bson:"phone"`
	Password     string             `bson:"password"`
	Address      string             `bson:"address"`
	SearchString string             `bson:"searchstring"`
	ResourceID   primitive.ObjectID `bson:"resourceid"`
	SmallImage   string             `bson:"smallimage"`
	MediumImage  string             `bson:"mediumimage"`
	LargeImage   string             `bson:"largeimage"`
}

// CustomerUpdateBSON ...
type CustomerUpdateBSON struct {
	Name         string             `bson:"name"`
	Password     string             `bson:"password"`
	Address      string             `bson:"address"`
	SearchString string             `bson:"searchstring"`
	ResourceID   primitive.ObjectID `bson:"resourceid"`
	SmallImage   string             `bson:"smallimage"`
	MediumImage  string             `bson:"mediumimage"`
	LargeImage   string             `bson:"largeimage"`
}

// ConvertToCreateBSON ...
func (payload CustomerRegisterPayload) ConvertToCreateBSON() (customerBSON CustomerCreateBSON) {
	// Get default img name
	img := config.GetIMG()
	customerBSON = CustomerCreateBSON{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Phone:        payload.Phone,
		Password:     util.Hash(payload.Password),
		Address:      payload.Address,
		SearchString: util.ConvertToHex(payload.Name),
		SmallImage:   img.SmallImage,
		MediumImage:  img.MediumImage,
		LargeImage:   img.LargeImage,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload CustomerUpdatePayload) ConvertToUpdateBSON(resource Resource) (customerBSON CustomerUpdateBSON) {
	customerBSON = CustomerUpdateBSON{
		Name:         payload.Name,
		Password:     payload.Password,
		Address:      payload.Address,
		SearchString: util.ConvertToHex(payload.Name),
		ResourceID:   payload.ResourceObjectID,
		SmallImage:   resource.SmallImage.GetName(),
		MediumImage:  resource.MediumImage.GetName(),
		LargeImage:   resource.LargeImage.GetName(),
	}
	return
}

// GenerateToken ...
func (c Customer) GenerateToken() (token string, err error) {
	env := config.GetENV()
	token, err = util.TokenEncode(c.ID.Hex(), env.CustomerKey)
	return
}

// ConvertToUpdateSearchStringBSON ...
func (c Customer) ConvertToUpdateSearchStringBSON() (companyBSON CustomerUpdateBSON) {
	companyBSON = CustomerUpdateBSON{
		Name:         c.Name,
		Password:     util.Hash(c.Password),
		Address:      c.Address,
		SearchString: util.ConvertToHex(c.Name),
	}
	return
}
