package user

import (
	controller "garagesvc/controller/user"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userCompany(g *echo.Group) {
	group := g.Group("/companies")

	group.GET("", controller.CompanyList, validator.CompanyFindRequest)
	group.GET("/:id", controller.CompanyDetail, validator.CompanyCheckExistance)
}
