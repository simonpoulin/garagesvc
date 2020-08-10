package models

type Role struct {
	ID   int    `json:"_id" bson:"_id"`
	Role string `json:"role" bson:"status"`
}
