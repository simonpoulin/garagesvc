package admin

import (
	controller "garagesvc/controller/admin"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminCompany(g *echo.Group) {
	group := g.Group("/companies")

	group.GET("/", controller.CompanyList, validator.CompanyFindRequest)
	group.GET("/:id", controller.CompanyDetail, validator.CompanyCheckExistance)
	group.POST("/", controller.CompanyCreate, validator.CompanyCreate)
	group.PATCH("/:id", controller.CompanyUpdate, validator.CompanyCheckExistance, validator.CompanyUpdate)
	group.DELETE("/:id", controller.CompanyDelete, validator.CompanyCheckExistance)
}
