package service

import (
	"errors"
	"fmt"
	"garagesvc/config"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResourceUpload ...
func ResourceUpload(file *multipart.FileHeader) (resourceID primitive.ObjectID, err error) {

	var (
		resource  model.Resource
		largeImg  model.Image
		mediumImg model.Image
		smallImg  model.Image
	)

	// Open multipart file
	src, err := file.Open()
	if err != nil {
		return
	}

	// Get file's extension
	nameParts := strings.Split(file.Filename, ".")
	extension := strings.ToLower("." + nameParts[len(nameParts)-1])

	// Check if file's extension is accepted
	ext := config.GetEXT()
	found := util.Find(ext.Extensions, extension)
	if !found {
		err = errors.New("file not allowed")
		return
	}

	// Get image directory
	env := config.GetENV()
	imgDir := env.ImageDirectory + "/"

	// Set large image destination
	imgPath := imgDir + primitive.NewObjectID().Hex() + extension

	// Create new file at destination
	dst, err := os.Create(imgPath)
	if err != nil {
		return
	}

	// Copy multipart file to file at destination
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	src.Close()
	dst.Close()

	// Get image
	image, err := imaging.Open(imgPath)
	if err != nil {
		return
	}

	// Set large image infomation
	largeImg.ID = primitive.NewObjectID()
	largeImg.Size.Height = 200
	largeImg.Size.Width = 200
	largeImg.Extension = extension

	// Resize the image to width = 200px, height = 200px
	largeImage := imaging.Resize(image, 200, 200, imaging.Lanczos)

	// Set large image destination
	largeImgPath := imgDir + largeImg.ID.Hex() + extension

	// Save the resulting image
	err = imaging.Save(largeImage, largeImgPath)
	if err != nil {
		return
	}

	// Set medium image infomation
	mediumImg.ID = primitive.NewObjectID()
	mediumImg.Size.Height = 120
	mediumImg.Size.Width = 120
	mediumImg.Extension = extension

	// Resize the image to width = 120px, height = 120px
	mediumImage := imaging.Resize(image, 120, 120, imaging.Lanczos)

	// Set medium image destination
	mediumImgPath := imgDir + mediumImg.ID.Hex() + extension

	// Save the resulting image
	err = imaging.Save(mediumImage, mediumImgPath)
	if err != nil {
		return
	}

	// Set small image infomation
	smallImg.ID = primitive.NewObjectID()
	smallImg.Size.Height = 32
	smallImg.Size.Width = 32
	smallImg.Extension = extension

	// Resize the image to width = 32px, height = 32px
	smallImage := imaging.Resize(image, smallImg.Size.Width, smallImg.Size.Height, imaging.Lanczos)

	// Set small image destination
	smallImgPath := imgDir + smallImg.ID.Hex() + extension

	// Save the resulting image
	err = imaging.Save(smallImage, smallImgPath)
	if err != nil {
		return
	}

	// Set medium file data
	resource.ID = primitive.NewObjectID()
	resource.SmallImage = smallImg
	resource.MediumImage = mediumImg
	resource.LargeImage = largeImg

	// Add to DB
	err = dao.ResourceCreate(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Delete source image
	err = os.Remove(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Return resource ID
	resourceID = resource.ID
	return
}

// ResourceDelete ...
func ResourceDelete(id primitive.ObjectID) (err error) {

	//Set filter
	filter := bson.M{"_id": id}

	//Get resources to delete images
	resource, err := dao.ResourceFindOne(filter)
	if err != nil {
		return
	}

	//Delete images from resource
	err = resource.DeleteImages()
	if err != nil {
		return
	}

	//Delete employee
	err = dao.ResourceDelete(filter)
	return
}
