package model

import (
	"garagesvc/config"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer ...
type Customer struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name" validator:"required, max=10, alphanumunicode"`
	Phone    string             `json:"phone" bson:"phone" validator:"required, eq=10, numeric"`
	Password string             `json:"password" bson:"password" validator:"required, min=6, max=20"`
}

// CustomerLoginPayload ...
type CustomerLoginPayload struct {
	Phone    string `json:"phone" bson:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" bson:"password" valid:"required, type(string), stringlength(6|20)"`
}

// CustomerRegisterPayload ...
type CustomerRegisterPayload struct {
	Name     string `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Phone    string `json:"phone" bson:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" bson:"password" valid:"required, type(string), stringlength(6|20)"`
}

// CustomerUpdatePayload ...
type CustomerUpdatePayload struct {
	Name     string `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Password string `json:"password" bson:"password" valid:"required, type(string), stringlength(6|20)"`
}

// GenerateToken ...
func (c Customer) GenerateToken() (token string, err error) {
	env := config.GetENV()
	token, err = util.TokenEncode(c.ID.Hex(), env.CustomerKey)
	return
}
