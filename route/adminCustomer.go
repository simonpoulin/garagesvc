package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminCustomer(g *echo.Group) {
	group := g.Group("/customers")

	group.GET("/", controller.CustomerList)
	group.GET("/:id", controller.CustomerDetail, validator.CustomerCheckExistance)
	group.PATCH("/:id", controller.CustomerUpdate, validator.CustomerCheckExistance, validator.CustomerValid)
	group.DELETE("/:id", controller.CustomerDelete, validator.CustomerCheckExistance)
}
