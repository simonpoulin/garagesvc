package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Image ...
type Image struct {
	ID        primitive.ObjectID `bson:"_id"`
	Size      Size               `bson:"size"`
	Extension string             `bson:"extension"`
}

// GetName ...
func (img Image) GetName() string {
	return img.ID.Hex() + img.Extension
}
