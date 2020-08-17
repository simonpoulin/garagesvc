package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo"
)

// BookingCreate ...
func BookingCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.BookingCreatePayload)
	)

	//Create booking
	result, err := service.BookingCreate(payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingDetail ...
func BookingDetail(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Get booking by ID
	result, err := service.BookingDetail(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
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

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingListByServiceID ...
func BookingListByServiceID(c echo.Context) error {
	var (
		serviceID = c.Param("serviceid")
	)

	//Get booking list
	result, err := service.BookingListByServiceID(serviceID)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingListByCustomerID ...
func BookingListByCustomerID(c echo.Context) error {
	var (
		customerID = c.Param("customerid")
	)

	//Get booking list
	result, err := service.BookingListByCustomerID(customerID)

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
		id      = c.Param("id")
		payload = c.Get("body").(model.BookingUpdatePayload)
	)

	//Update booking
	result, err := service.BookingUpdate(id, payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingChangeStatus ...
func BookingChangeStatus(c echo.Context) error {
	var (
		id     = c.Param("id")
		status = c.Param("status")
	)

	//Change booking active state
	result, err := service.BookingChangeStatus(id, status)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// BookingDelete ...
func BookingDelete(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Delete booking by ID
	err := service.BookingDelete(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
