package route

import (
	"garagesvc/auth"
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func booking(e *echo.Echo) {
	group := e.Group("/bookings")

	group.Use(auth.IsLoggedIn)

	group.GET("/", controller.BookingList)
	group.GET("/:id", controller.BookingDetail)
	group.GET("/status/:status", controller.BookingListByStatus)
	group.GET("/service/:serviceid", controller.BookingListByServiceID)
	group.GET("/customer/:customerid", controller.BookingListByCustomerID)
	group.POST("/", controller.BookingCreate, validator.BookingCreate)
	group.PATCH("/:id", controller.BookingUpdate, validator.BookingUpdate)
	group.PATCH("/:id/status/:status", controller.BookingChangeStatus)
	group.DELETE("/:id", controller.BookingDelete)
}
