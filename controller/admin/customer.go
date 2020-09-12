package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

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
	var (
		name = c.QueryParam("name")
		page = c.Get("page").(int)
	)

	//Get customer list
	result, err := service.CustomerList(name, page)

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
		payload  = c.Get("body").(model.CustomerUpdatePayload)
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
