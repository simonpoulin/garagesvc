package route

import (
	"garagesvc/controller"
	"garagesvc/validator"

	"github.com/labstack/echo"
)

func booking(e *echo.Echo) {
	// Customer
	customerGroup := e.Group("/:customerID/bookings")

	customerGroup.GET("/", controller.BookingList)
	customerGroup.GET("/:id", controller.BookingDetail)
	customerGroup.GET("/:id/status/:status", controller.BookingListByStatus)
	customerGroup.POST("/", controller.BookingCreate)
	customerGroup.PATCH("/:id", controller.BookingUpdate)
	customerGroup.DELETE("/:id", controller.BookingDelete)

	// Service
	serviceGroup := e.Group("/companies/:companyID/services/:serviceID/bookings")

	serviceGroup.Use(validator.ServiceCheck, validator.CompanyCheck)

	serviceGroup.GET("/", controller.BookingList)
	serviceGroup.GET("/:id", controller.BookingDetail)
	serviceGroup.GET("/status/:status", controller.BookingListByStatus)
	serviceGroup.POST("/", controller.BookingCreate)
	serviceGroup.PATCH("/:id", controller.BookingUpdate)
	serviceGroup.PATCH("/:id/status/:status", controller.BookingChangeStatus)
	serviceGroup.DELETE("/:id", controller.BookingDelete)
}
