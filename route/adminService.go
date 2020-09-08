package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminService(g *echo.Group) {
	group := g.Group("/services")

	group.GET("/", controller.ServiceList)
	group.GET("/:id", controller.ServiceDetail, validator.ServiceCheckExistance)
	group.GET("/active/:active", controller.ServiceListByActiveState)
	group.GET("/company/:companyid", controller.ServiceListByCompanyID)
	group.POST("/", controller.ServiceCreate, validator.ServiceCreate)
	group.PATCH("/:id", controller.ServiceUpdate, validator.ServiceCheckExistance, validator.ServiceUpdate)
	group.PATCH("/:id/active", controller.ServiceChangeActive, validator.ServiceCheckExistance)
	group.DELETE("/:id", controller.ServiceDelete, validator.ServiceCheckExistance)
}
