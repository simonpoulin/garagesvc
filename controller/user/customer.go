package user

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
