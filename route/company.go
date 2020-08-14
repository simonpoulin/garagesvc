package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func company(e *echo.Echo) {
	group := e.Group("/companies")

	group.GET("/", controller.CompanyList)
	group.GET("/:id", controller.CompanyDetail)
	group.POST("/", controller.CompanyCreate, validator.CompanyCreate)
	group.PATCH("/:id", controller.CompanyUpdate, validator.CompanyUpdate)
	group.PATCH("/:id/active", controller.CompanyChangeActive)
	group.DELETE("/:id", controller.CompanyDelete)
}
