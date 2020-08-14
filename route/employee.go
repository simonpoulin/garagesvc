package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func employee(e *echo.Echo) {
	group := e.Group("/employees")

	group.GET("/", controller.EmployeeList)
	group.GET("/:id", controller.EmployeeDetail)
	group.GET("/active/:active", controller.EmployeeListByActiveState)
	group.POST("/", controller.EmployeeCreate, validator.EmployeeCreate)
	group.POST("/login", controller.EmployeeLogin, validator.EmployeeLogin)
	group.PATCH("/:id", controller.EmployeeUpdate, validator.EmployeeUpdate)
	group.PATCH("/:id/active", controller.EmployeeChangeActive)
	group.DELETE("/:id", controller.EmployeeDelete)
}
