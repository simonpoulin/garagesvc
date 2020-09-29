package user

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// BookingCreate godoc
//
// @Summary User API - Booking create
// @Description Create a booking
// @Tags User - Bookings
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
// @Router /user/bookings [post]
func BookingCreate(c echo.Context) error {
	var (
		payload  = c.Get("body").(model.BookingCreatePayload)
		customer = c.Get("authcustomer").(model.Customer)
	)

	//Create booking
	result, err := service.BookingCreate(payload, customer.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingDetail godoc
//
// @Summary User API - Booking detail
// @Description Return details of a booking
// @Tags User - Bookings
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
// @Router /user/bookings/{id} [get]
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
// @Summary User API - Booking list
// @Description Return a list of bookings
// @Tags User - Bookings
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
// @Router /user/bookings [get]
func BookingList(c echo.Context) error {
	var (
		queryValues = c.Get("query").(model.BookingQuery)
		query       = model.AppQuery{
			Status:     queryValues.Status,
			Page:       queryValues.Page,
			ServiceID:  queryValues.ServiceObjectID,
			CustomerID: queryValues.CustomerObjectID,
		}
	)

	//Get booking list
	result, err := service.BookingList(query)

	//Handle errors
	if err != nil {
		//If list is not empty, return 400
		if !util.IsEmptyListError(err) {
			return util.Response400(c, err.Error())
		}
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingUpdate godoc
//
// @Summary User API - Booking update
// @Description Update booking's details
// @Tags User - Bookings
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
// @Router /user/bookings/{id} [patch]
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
// @Summary User API - Booking delete
// @Description Delete a booking
// @Tags User - Bookings
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
// @Router /user/bookings/{id} [delete]
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
