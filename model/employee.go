package model

import (
	"garagesvc/config"
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee ...
type Employee struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Phone    string             `bson:"phone"`
	Password string             `bson:"password"`
	Active   bool               `bson:"active"`
}

// EmployeeCreateBSON ...
type EmployeeCreateBSON struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Phone        string             `bson:"phone"`
	Password     string             `bson:"password"`
	Active       bool               `bson:"active"`
	SearchString string             `bson:"searchstring"`
}

// EmployeeUpdateBSON ...
type EmployeeUpdateBSON struct {
	Name         string `bson:"name"`
	Phone        string `bson:"phone"`
	Password     string `bson:"password"`
	Active       bool   `bson:"active"`
	SearchString string `bson:"searchstring"`
}

// ConvertToCreateBSON ...
func (payload EmployeeRegisterPayload) ConvertToCreateBSON() (employeeBSON EmployeeCreateBSON) {
	employeeBSON = EmployeeCreateBSON{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Phone:        payload.Phone,
		Password:     util.Hash(payload.Password),
		SearchString: util.ConvertToHex(payload.Name),
		Active:       true,
	}
	return
}

// ConvertToUpdateBSON ...
func (payload EmployeeUpdatePayload) ConvertToUpdateBSON() (employeeBSON EmployeeUpdateBSON) {
	employeeBSON = EmployeeUpdateBSON{
		Name:         payload.Name,
		Phone:        payload.Phone,
		Password:     payload.Password,
		Active:       payload.Active,
		SearchString: util.ConvertToHex(payload.Name),
	}
	return
}

// GenerateToken ...
func (e Employee) GenerateToken() (token string, err error) {
	env := config.GetENV()
	token, err = util.TokenEncode(e.ID.Hex(), env.EmployeeKey)
	return
}

// ConvertToUpdateSearchStringBSON ...
func (e Employee) ConvertToUpdateSearchStringBSON() (employeeBSON EmployeeUpdateBSON) {
	employeeBSON = EmployeeUpdateBSON{
		Name:         e.Name,
		Phone:        e.Phone,
		Password:     util.Hash(e.Password),
		Active:       e.Active,
		SearchString: util.ConvertToHex(e.Name),
	}
	return
}
