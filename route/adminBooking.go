package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func adminBooking(g *echo.Group) {
	group := g.Group("/bookings")

	group.GET("/", controller.BookingList)
	group.GET("/:id", controller.BookingDetail, validator.BookingCheckExistance)
	group.GET("/status/:status", controller.BookingListByStatus)
	group.GET("/service/:serviceid", controller.BookingListByServiceID, validator.ServiceCheckExistance)
	group.GET("/customer/:customerid", controller.BookingListByCustomerID, validator.CustomerCheckExistance)
	group.POST("/", controller.BookingCreate, validator.BookingCreate)
	group.PATCH("/:id", controller.BookingUpdate, validator.BookingCheckExistance, validator.BookingUpdate)
	group.DELETE("/:id", controller.BookingDelete, validator.BookingCheckExistance)
}
