package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo"
)

// EmployeeCreate ...
func EmployeeCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.EmployeeCreatePayload)
	)

	//Create employee
	result, err := service.EmployeeCreate(payload, c)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeLogin ...
func EmployeeLogin(c echo.Context) error {
	var (
		payload = c.Get("body").(model.EmployeeLoginPayload)
	)

	//Create token
	result, err := service.EmployeeLogin(payload, c)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error(), nil)
	}

	//Return 200
	return util.Response200(c, "", result)
}
