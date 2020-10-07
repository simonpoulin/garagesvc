package route

import (
	"garagesvc/config"
	controller "garagesvc/controller/common"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func resource(e *echo.Echo) {
	group := e.Group("resources")

	env := config.GetENV()
	group.DELETE("/{id}", controller.ResourceDelete, validator.ResourceCheckExistance)
	group.POST("/upload", controller.ResourceUpload, validator.ResourceUpload)
	group.Static("/img", env.ImageDirectory)
}
