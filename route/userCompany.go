package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userCompany(g *echo.Group) {
	group := g.Group("/companies")

	group.GET("/", controller.CompanyList)
	group.GET("/:id", controller.CompanyDetail, validator.CompanyCheckExistance)
}
