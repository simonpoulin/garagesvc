package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// BookingCreate ...
func BookingCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.BookingCreatePayload)
	)

	//Create booking
	result, err := service.BookingCreate(payload)

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

	//Get booking list
	result, err := service.BookingList()

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingListByStatus ...
func BookingListByStatus(c echo.Context) error {
	var (
		status = c.Param("status")
	)

	//Get booking list
	result, err := service.BookingListByStatus(status)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingListByServiceID ...
func BookingListByServiceID(c echo.Context) error {
	var (
		svc = c.Get("service").(model.Service)
	)

	//Get booking list
	result, err := service.BookingListByServiceID(svc.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingListByCustomerID ...
func BookingListByCustomerID(c echo.Context) error {
	var (
		customer = c.Get("customer").(model.Customer)
	)

	//Get booking list
	result, err := service.BookingListByCustomerID(customer.ID)

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
