package common

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// EmployeeRegister godoc
//
// @Summary Admin API - Employee Register
// @Description Create an employee
// @Tags Common
//
// @Accept  json
// @Produce  json
//
// @Param EmployeeRegisterPayload body model.EmployeeRegisterPayload true "Employee Register Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Router /employee-register [post]
func EmployeeRegister(c echo.Context) error {
	var (
		payload = c.Get("body").(model.EmployeeRegisterPayload)
	)

	//Create employee
	result, err := service.EmployeeRegister(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeLogin godoc
//
// @Summary Admin API - Employee Login
// @Description Return a token for logging employee
// @Tags Common
//
// @Accept  json
// @Produce  json
//
// @Param EmployeeLoginPayload body model.EmployeeLoginPayload true "Employee Login Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Router /employee-login [post]
func EmployeeLogin(c echo.Context) error {
	var (
		payload = c.Get("body").(model.EmployeeLoginPayload)
	)

	//Create token
	result, err := service.EmployeeLogin(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}
