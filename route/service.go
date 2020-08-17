package route

import (
	"garagesvc/auth"
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func service(e *echo.Echo) {
	group := e.Group("/services")

	group.Use(auth.IsLoggedIn)

	group.GET("/", controller.ServiceList)
	group.GET("/:id", controller.ServiceDetail)
	group.GET("/status/:status", controller.ServiceListByActiveState)
	group.GET("/company/:companyid", controller.ServiceListByCompanyID)
	group.POST("/", controller.ServiceCreate, validator.ServiceCreate)
	group.PATCH("/:id", controller.ServiceUpdate, validator.ServiceUpdate)
	group.PATCH("/:id/active", controller.ServiceChangeActive)
	group.DELETE("/:id", controller.ServiceDelete)
}
