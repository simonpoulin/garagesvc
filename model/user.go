package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User ...
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name" validator:"required, max=10, alphanumunicode"`
	Phone    string             `json:"phone" bson:"phone" validator:"required, eq=10, numeric"`
	Password string             `json:"password" bson:"password" validator:"required, min=6, max=20"`
}
