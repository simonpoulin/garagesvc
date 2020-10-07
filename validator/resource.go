package validator

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResourceUpload ...
func ResourceUpload(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return util.Response400(c, err.Error())
		}
		c.Set("file", file)
		return next(c)
	}
}

// ResourceCheckExistance ...
func ResourceCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		//Validate resource
		resource, err := ResourceValidate(id)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set body and move to next process
		c.Set("resource", resource)
		return next(c)
	}
}

// ResourceValidate ...
func ResourceValidate(resourceID string) (resource model.Resource, err error) {

	//Check valid resource ID
	rscID, err := primitive.ObjectIDFromHex(resourceID)
	if err != nil {
		return
	}

	//Set filter
	filter := bson.M{"_id": rscID}

	//Validate resource
	resource, err = dao.ResourceFindOne(filter)

	return
}
