package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// ServiceResponse ...
type ServiceResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	Company     CompanyResponse    `json:"company"`
	Name        string             `json:"name"`
	Location    Location           `json:"location"`
	Address     string             `json:"address"`
	Active      bool               `json:"active"`
	Phone       string             `json:"phone"`
	Email       string             `json:"email"`
	Description string             `json:"description"`
	ResourceID  primitive.ObjectID `json:"resourceid"`
	SmallImage  string             `json:"smallimage"`
	MediumImage string             `json:"mediumimage"`
	LargeImage  string             `json:"largeimage"`
}
