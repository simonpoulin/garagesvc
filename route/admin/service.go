package admin

import (
	controller "garagesvc/controller/admin"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminService(g *echo.Group) {
	group := g.Group("/services")

	group.GET("/", controller.ServiceList, validator.ServiceFindRequest)
	group.GET("/:id", controller.ServiceDetail, validator.ServiceCheckExistance)
	group.POST("/", controller.ServiceCreate, validator.ServiceCreate)
	group.PATCH("/:id", controller.ServiceUpdate, validator.ServiceCheckExistance, validator.ServiceUpdate)
	group.DELETE("/:id", controller.ServiceDelete, validator.ServiceCheckExistance)
}
