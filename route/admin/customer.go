package admin

import (
	controller "garagesvc/controller/admin"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminCustomer(g *echo.Group) {
	group := g.Group("/customers")

	group.GET("", controller.CustomerList, validator.CustomerFindRequest)
	group.GET("/:id", controller.CustomerDetail, validator.CustomerCheckExistance)
	group.PATCH("/:id", controller.CustomerUpdate, validator.CustomerCheckExistance, validator.CustomerUpdate)
	group.DELETE("/:id", controller.CustomerDelete, validator.CustomerCheckExistance)
}
