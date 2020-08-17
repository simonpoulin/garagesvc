package route

import (
	"garagesvc/auth"
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func customer(e *echo.Echo) {
	group := e.Group("/users")

	group.GET("/", controller.CustomerList, auth.IsLoggedIn)
	group.GET("/:id", controller.CustomerDetail, auth.IsLoggedIn)
	group.POST("/", controller.CustomerCreate, validator.CustomerValid)
	group.POST("/login", controller.CustomerLogin, validator.CustomerLogin)
	group.PATCH("/:id", controller.CustomerUpdate, validator.CustomerValid, auth.IsLoggedIn)
	group.DELETE("/:id", controller.CustomerDelete, auth.IsLoggedIn)
}
