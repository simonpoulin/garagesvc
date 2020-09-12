package user

import (
	controller "garagesvc/controller/user"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userCustomer(g *echo.Group) {
	group := g.Group("/customers")

	group.GET("/:id", controller.CustomerDetail, validator.CustomerOwner)
	group.PATCH("/:id", controller.CustomerUpdate, validator.CustomerOwner, validator.CustomerUpdate)
}
