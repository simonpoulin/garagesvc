package route

import (
	controller "garagesvc/controller/common"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func common(e *echo.Echo) {
	group := e.Group("")

	group.POST("/employee-register", controller.EmployeeCreate, validator.EmployeeCreate)
	group.POST("/employee-login", controller.EmployeeLogin, validator.EmployeeLogin)
	group.POST("/customer-register", controller.CustomerRegister, validator.CustomerCreate)
	group.POST("/customer-login", controller.CustomerLogin, validator.CustomerLogin)
}
