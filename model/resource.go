package model

import (
	"garagesvc/config"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Resource ...
type Resource struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	SmallImage  Image              `json:"smallimage" bson:"smallimage"`
	MediumImage Image              `json:"mediumimage" bson:"mediumimage"`
	LargeImage  Image              `json:"largeimage" bson:"largeimage"`
}

// DeleteImages ...
func (resource Resource) DeleteImages() (err error) {
	var (
		env = config.GetENV()
		dir = env.ImageDirectory
	)
	sImgPath := dir + "/" + resource.SmallImage.GetName()
	err = os.Remove(sImgPath)
	if err != nil {
		return
	}
	mImgPath := dir + "/" + resource.MediumImage.GetName()
	err = os.Remove(mImgPath)
	if err != nil {
		return
	}
	lImgPath := dir + "/" + resource.LargeImage.GetName()
	err = os.Remove(lImgPath)
	if err != nil {
		return
	}
	return
}

// GetDefaultResource ...
func (resource Resource) GetDefaultResource() {
	var (
		smallImg  Image
		mediumImg Image
		largeImg  Image
	)

	smallImg.ID, _ = primitive.ObjectIDFromHex("5f7691bee77eedd2166d4343")
	smallImg.Size.Height = 32
	smallImg.Size.Width = 32
	smallImg.Extension = ".png"
	resource.SmallImage = smallImg

	mediumImg.ID, _ = primitive.ObjectIDFromHex("5f7691bee77eedd2166d4342")
	mediumImg.Size.Height = 120
	mediumImg.Size.Width = 120
	mediumImg.Extension = ".png"
	resource.MediumImage = mediumImg

	largeImg.ID, _ = primitive.ObjectIDFromHex("5f7691bee77eedd2166d4341")
	largeImg.Size.Height = 200
	largeImg.Size.Width = 200
	largeImg.Extension = ".png"
	resource.LargeImage = largeImg
}
