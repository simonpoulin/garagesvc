package user

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CustomerDetail godoc
//
// @Summary User API - Customer detail
// @Description Return details of a customer
// @Tags User - Customers
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Customer's ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /user/customers [get]
func CustomerDetail(c echo.Context) error {
	var (
		customer = c.Get("authcustomer").(model.Customer)
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

// CustomerUpdate godoc
//
// @Summary User API - Customer update
// @Description Update customer's details
// @Tags User - Customers
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Customer's ID"
// @Param CustomerUpdatePayload body model.CustomerUpdatePayload true "Customer Update Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /user/customers [patch]
func CustomerUpdate(c echo.Context) error {
	var (
		customer = c.Get("authcustomer").(model.Customer)
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
