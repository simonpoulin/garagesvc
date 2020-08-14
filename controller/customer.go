package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo"
)

// CustomerCreate ...
func CustomerCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CustomerPayload)
	)

	//Create customer
	result, err := service.CustomerCreate(payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerLogin ...
func CustomerLogin(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CustomerLoginPayload)
	)

	//Create token
	result, err := service.CustomerLogin(payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerDetail ...
func CustomerDetail(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Get customer by ID
	result, err := service.CustomerDetail(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerList ...
func CustomerList(c echo.Context) error {

	//Get customer list
	result, err := service.CustomerList()

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerUpdate ...
func CustomerUpdate(c echo.Context) error {
	var (
		id      = c.Param("id")
		payload = c.Get("body").(model.CustomerPayload)
	)

	//Update customer
	result, err := service.CustomerUpdate(id, payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerDelete ...
func CustomerDelete(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Delete customer by ID
	err := service.CustomerDelete(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", nil)
}
