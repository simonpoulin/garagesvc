package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookingCreate godoc
//
// @Summary Admin API - Booking create
// @Description Create a booking
// @Tags Admin - Bookings
//
// @Accept  json
// @Produce  json
//
// @Param BookingCreatePayload body model.BookingCreatePayload true "Booking Create Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/bookings/ [post]
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

// BookingDetail godoc
//
// @Summary Admin API - Booking detail
// @Description Return details of a booking
// @Tags Admin - Bookings
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Booking ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/bookings/{id} [get]
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

// BookingList godoc
//
// @Summary Admin API - Booking list
// @Description Return a list of bookings
// @Tags Admin - Bookings
//
// @Accept  json
// @Produce  json
//
// @Param status query string false "Status state"
// @Param serviceid query string false "Service's ID"
// @Param customerid query string false "Customer's ID"
// @Param page query int false "Page number"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/bookings/ [get]
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

// BookingUpdate godoc
//
// @Summary Admin API - Booking update
// @Description Update booking's details
// @Tags Admin - Bookings
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Booking's ID"
// @Param BookingUpdatePayload body model.BookingUpdatePayload true "Booking Update Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/bookings/{id} [patch]
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

// BookingDelete godoc
//
// @Summary Admin API - Booking delete
// @Description Delete a booking
// @Tags Admin - Bookings
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Booking's ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/bookings/{id} [delete]
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
