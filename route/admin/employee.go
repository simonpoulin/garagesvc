package admin

import (
	controller "garagesvc/controller/admin"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminEmployee(g *echo.Group) {
	group := g.Group("/employees")

	group.GET("/", controller.EmployeeList, validator.EmployeeFindRequest)
	group.GET("/:id", controller.EmployeeDetail, validator.EmployeeCheckExistance)
	group.PATCH("/:id", controller.EmployeeUpdate, validator.EmployeeCheckExistance, validator.EmployeeUpdate)
	group.DELETE("/:id", controller.EmployeeDelete, validator.EmployeeCheckExistance)
}
