package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userBooking(g *echo.Group) {
	group := g.Group("/bookings")

	group.GET("/:id", controller.BookingDetail, validator.BookingCheckExistance)
	group.POST("/", controller.BookingCreate, validator.BookingCreate)
	group.PATCH("/:id", controller.BookingUpdate, validator.BookingCheckExistance, validator.BookingUpdate)
	group.DELETE("/:id", controller.BookingDelete, validator.BookingCheckExistance)
}
