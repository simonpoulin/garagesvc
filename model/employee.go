package model

import (
	"garagesvc/config"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee ...
type Employee struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Phone    string             `json:"phone" bson:"phone"`
	Password string             `json:"password" bson:"password"`
	Active   bool               `json:"active" bson:"active"`
}

// EmployeeLoginPayload ...
type EmployeeLoginPayload struct {
	Phone    string `json:"phone" bson:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" bson:"password" valid:"required, type(string), stringlength(6|20)"`
}

// EmployeeCreatePayload ...
type EmployeeCreatePayload struct {
	Name     string `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Phone    string `json:"phone" bson:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" bson:"password" valid:"required, type(string), stringlength(6|20)"`
}

// EmployeeUpdatePayload ...
type EmployeeUpdatePayload struct {
	Name     string `json:"name" bson:"name" valid:"required, stringlength(1|20)"`
	Phone    string `json:"phone" bson:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" bson:"password" valid:"required, type(string), stringlength(6|20)"`
	Active   bool   `json:"active" bson:"active" valid:"required, type(bool)"`
}

// GenerateToken ...
func (e Employee) GenerateToken() (token string, err error) {
	env := config.GetENV()
	token, err = util.TokenEncode(e.ID.Hex(), env.EmployeeKey)
	return
}
