package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreate ...
func BookingCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.BookingCreatePayload)
	)

	//Create booking
	result, err := service.BookingCreate(payload, nil)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingDetail ...
func BookingDetail(c echo.Context) error {
	var (
		booking = c.Get("booking").(model.Booking)
	)

	//Get booking by ID
	result, err := service.BookingDetail(booking.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingList ...
func BookingList(c echo.Context) error {
	var (
		status     = c.QueryParam("status")
		serviceID  = c.Get("serviceID").(primitive.ObjectID)
		customerID = c.Get("customerID").(primitive.ObjectID)
		page       = c.Get("page").(int)
	)

	//Get booking list
	result, err := service.BookingList(status, serviceID, customerID, page)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingUpdate ...
func BookingUpdate(c echo.Context) error {
	var (
		booking = c.Get("booking").(model.Booking)
		payload = c.Get("body").(model.BookingUpdatePayload)
	)

	//Update booking
	result, err := service.BookingUpdate(booking.ID, payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingDelete ...
func BookingDelete(c echo.Context) error {
	var (
		booking = c.Get("booking").(model.Booking)
	)

	//Delete booking by ID
	err := service.BookingDelete(booking.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
