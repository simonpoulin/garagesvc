package user

import (
	controller "garagesvc/controller/user"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userCustomer(g *echo.Group) {
	group := g.Group("/customers")

	group.GET("", controller.CustomerDetail)
	group.PATCH("", controller.CustomerUpdate, validator.CustomerUpdate)
}
