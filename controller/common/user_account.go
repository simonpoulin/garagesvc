package common

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CustomerRegister ...
func CustomerRegister(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CustomerCreatePayload)
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
