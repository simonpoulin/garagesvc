package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// EmployeeResponse ...
type EmployeeResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Phone    string             `json:"phone"`
	Password string             `json:"password"`
	Active   bool               `json:"active"`
}
