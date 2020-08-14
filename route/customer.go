package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func customer(e *echo.Echo) {
	group := e.Group("/users")

	group.GET("/", controller.CustomerList)
	group.GET("/:id", controller.CustomerDetail)
	group.POST("/", controller.CustomerCreate, validator.CustomerValid)
	group.POST("/login", controller.CustomerLogin, validator.CustomerLogin)
	group.PATCH("/:id", controller.CustomerUpdate, validator.CustomerValid)
	group.DELETE("/:id", controller.CustomerDelete)
}
