package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminEmployee(g *echo.Group) {
	group := g.Group("/employees")

	group.GET("/", controller.EmployeeList)
	group.GET("/:id", controller.EmployeeDetail, validator.EmployeeCheckExistance)
	group.GET("/active/:active", controller.EmployeeListByActiveState)
	group.PATCH("/:id", controller.EmployeeUpdate, validator.EmployeeCheckExistance, validator.EmployeeUpdate)
	group.PATCH("/:id/active", controller.EmployeeChangeActive, validator.EmployeeCheckExistance)
	group.DELETE("/:id", controller.EmployeeDelete, validator.EmployeeCheckExistance)
}
