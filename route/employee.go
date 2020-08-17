package route

import (
	"garagesvc/auth"
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func employee(e *echo.Echo) {
	group := e.Group("/employees")

	group.GET("/", controller.EmployeeList, auth.IsLoggedIn)
	group.GET("/:id", controller.EmployeeDetail, auth.IsLoggedIn)
	group.GET("/active/:active", controller.EmployeeListByActiveState, auth.IsLoggedIn)
	group.POST("/", controller.EmployeeCreate, validator.EmployeeCreate)
	group.POST("/login", controller.EmployeeLogin, validator.EmployeeLogin)
	group.PATCH("/:id", controller.EmployeeUpdate, validator.EmployeeUpdate, auth.IsLoggedIn)
	group.PATCH("/:id/active", controller.EmployeeChangeActive, auth.IsLoggedIn)
	group.DELETE("/:id", controller.EmployeeDelete, auth.IsLoggedIn)
}
