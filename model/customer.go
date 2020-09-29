package model

import (
	"garagesvc/config"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer ...
type Customer struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Phone    string             `json:"phone" bson:"phone"`
	Password string             `json:"password" bson:"password"`
}

// CustomerCreateBSON ...
type CustomerCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Phone        string             `bson:"phone"`
	Password     string             `bson:"password"`
	SearchString string             `bson:"searchstring"`
}

// CustomerUpdateBSON ...
type CustomerUpdateBSON struct {
	Name         string `bson:"name"`
	Password     string `bson:"password"`
	SearchString string `bson:"searchstring"`
}

// ConvertToCreateBSON ...
func (payload CustomerRegisterPayload) ConvertToCreateBSON() (customerBSON CustomerCreateBSON) {
	customerBSON = CustomerCreateBSON{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Phone:        payload.Phone,
		Password:     util.Hash(payload.Password),
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// ConvertToUpdateBSON ...
func (payload CustomerUpdatePayload) ConvertToUpdateBSON() (customerBSON CustomerUpdateBSON) {
	customerBSON = CustomerUpdateBSON{
		Name:         payload.Name,
		Password:     util.Hash(payload.Password),
		SearchString: util.ConvertToHex(payload.Name),
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
		SearchString: util.ConvertToHex(c.Name),
	}
	return
}
