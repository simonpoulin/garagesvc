package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Role     Role               `json:"role_id" bson:"role_id"`
	Name     string             `json:"name" bson:"name" validator:"required, max=10, alphanumunicode"`
	Phone    string             `json:"phone" bson:"phone" validator:"required, eq=10, numeric"`
	Password string             `json:"password" bson:"password" validator:"required, min=6, max=20"`
	Active   bool               `json:"active" bson:"active"`
}
