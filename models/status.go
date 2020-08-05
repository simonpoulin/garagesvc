package models

type Status struct {
	ID     int    `json:"_id" bson:"_id"`
	Status string `json:"status" bson:"status"`
}
