package user

import (
	controller "garagesvc/controller/user"
	"garagesvc/validator"

	"github.com/labstack/echo/v4"
)

func userBooking(g *echo.Group) {
	group := g.Group("/bookings")

	group.GET("", controller.BookingList, validator.BookingFindRequest)
	group.GET("/:id", controller.BookingDetail, validator.BookingOwner)
	group.POST("", controller.BookingCreate, validator.BookingCreate)
	group.PATCH("/:id", controller.BookingUpdate, validator.BookingOwner, validator.BookingUpdate)
	group.DELETE("/:id", controller.BookingDelete, validator.BookingOwner)
}
