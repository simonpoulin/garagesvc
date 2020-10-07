package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CustomerResponse ...
type CustomerResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Phone       string             `json:"phone"`
	Password    string             `json:"password"`
	Address     string             `json:"address"`
	ResourceID  primitive.ObjectID `json:"resourceid"`
	SmallImage  string             `json:"smallimage"`
	MediumImage string             `json:"mediumimage"`
	LargeImage  string             `json:"largeimage"`
}
