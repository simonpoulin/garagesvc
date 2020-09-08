package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CustomerCreate ...
func CustomerCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CustomerPayload)
	)

	//Create customer
	result, err := service.CustomerCreate(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
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

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerDetail ...
func CustomerDetail(c echo.Context) error {
	var (
		customer = c.Get("customer").(model.Customer)
	)

	//Get customer by ID
	result, err := service.CustomerDetail(customer.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerList ...
func CustomerList(c echo.Context) error {

	//Get customer list
	result, err := service.CustomerList()

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerUpdate ...
func CustomerUpdate(c echo.Context) error {
	var (
		customer = c.Get("customer").(model.Customer)
		payload  = c.Get("body").(model.CustomerPayload)
	)

	//Update customer
	result, err := service.CustomerUpdate(customer.ID, payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerDelete ...
func CustomerDelete(c echo.Context) error {
	var (
		customer = c.Get("customer").(model.Customer)
	)

	//Delete customer by ID
	err := service.CustomerDelete(customer.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
