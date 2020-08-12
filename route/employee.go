package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func employee(e *echo.Echo) {
	group := e.Group("/employees")

	// e.GET("/employees", GetEmployees)
	// e.GET("/employees/:id", GetEmployee)
	group.POST("/", controller.EmployeeCreate, validator.EmployeeCreate)
	group.POST("/login", controller.EmployeeLogin, validator.EmployeeLogin)
	// e.PATCH("/employees/:id", UpdateEmployee)
	// e.DELETE("/employees/:id", DeleteEmployee)
}
