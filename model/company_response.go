package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CompanyResponse ...
type CompanyResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Location Location           `json:"location"`
	Email    string             `json:"email"`
	Address  string             `json:"address"`
	Phone    string             `json:"phone"`
	Active   bool               `json:"active"`
}
