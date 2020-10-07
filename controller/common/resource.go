package common

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

//ResourceUpload ...
func ResourceUpload(c echo.Context) error {
	var (
		file = c.Get("file").(*multipart.FileHeader)
	)

	result, err := service.ResourceUpload(file)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)

}

// ResourceDelete...
func ResourceDelete(c echo.Context) error {
	var (
		resource = c.Get("resource").(model.Resource)
	)

	//Delete resource by ID
	err := service.ResourceDelete(resource.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
