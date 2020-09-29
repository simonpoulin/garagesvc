package user

import (
	controller "garagesvc/controller/user"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userService(g *echo.Group) {
	group := g.Group("/services")

	group.GET("/:id", controller.ServiceDetail, validator.ServiceCheckExistance)
	group.GET("", controller.ServiceList, validator.ServiceFindRequest)
}
