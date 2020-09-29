package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CustomerDetail godoc
//
// @Summary Admin API - Customer detail
// @Description Return details of a customer
// @Tags Admin - Customers
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
// @Router /admin/customers/{id} [get]
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

// CustomerList godoc
//
// @Summary Admin API - Customer list
// @Description Return a list of customers
// @Tags Admin - Customers
//
// @Accept  json
// @Produce  json
//
// @Param name query string false "Name keyword"
// @Param page query int false "Page number"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/customers [get]
func CustomerList(c echo.Context) error {
	var (
		queryValues = c.Get("query").(model.CustomerQuery)
		query       = model.AppQuery{
			Page: queryValues.Page,
			Name: queryValues.Name,
		}
	)

	//Get customer list
	result, err := service.CustomerList(query)

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

// CustomerUpdate godoc
//
// @Summary Admin API - Customer update
// @Description Update customer's details
// @Tags Admin - Customers
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
// @Router /admin/customers/{id} [patch]
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

// CustomerDelete godoc
//
// @Summary Admin API - Customer delete
// @Description Delete a customer
// @Tags Admin - Customers
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
// @Router /admin/customers/{id} [delete]
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
