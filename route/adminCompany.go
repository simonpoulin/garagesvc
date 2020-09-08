package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminCompany(g *echo.Group) {
	group := g.Group("/companies")

	group.GET("/", controller.CompanyList)
	group.GET("/:id", controller.CompanyDetail, validator.CompanyCheckExistance)
	group.POST("/", controller.CompanyCreate, validator.CompanyCreate)
	group.PATCH("/:id", controller.CompanyUpdate, validator.CompanyCheckExistance, validator.CompanyUpdate)
	group.PATCH("/:id/active", controller.CompanyChangeActive, validator.CompanyCheckExistance)
	group.DELETE("/:id", controller.CompanyDelete)
}
