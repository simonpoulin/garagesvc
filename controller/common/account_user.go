package common

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CustomerRegister godoc
//
// @Summary User API - Customer Register
// @Description Create an customer
// @Tags Common
//
// @Accept  json
// @Produce  json
//
// @Param CustomerRegisterPayload body model.CustomerRegisterPayload true "Customer Register Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Router /customer-register [post]
func CustomerRegister(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CustomerRegisterPayload)
	)

	//Create customer
	result, err := service.CustomerRegister(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CustomerLogin godoc
//
// @Summary User API - Customer Login
// @Description Return a token for logging customer
// @Tags Common
//
// @Accept  application/json
// @Produce  json
//
// @Param CustomerLoginPayload body model.CustomerLoginPayload true "Customer Login Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Router /customer-login [post]
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
