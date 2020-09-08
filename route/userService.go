package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userService(g *echo.Group) {
	group := g.Group("/services")

	group.GET("/:id", controller.ServiceDetail, validator.ServiceCheckExistance)
	group.GET("/active/:active", controller.ServiceListByActiveState)
	group.GET("/company/:companyid", controller.ServiceListByCompanyID)
}
