package common

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// EmployeeCreate ...
func EmployeeCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.EmployeeCreatePayload)
	)

	//Create employee
	result, err := service.EmployeeCreate(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
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
	result, err := service.EmployeeLogin(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}
