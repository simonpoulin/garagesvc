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
	result, err := service.EmployeeCreate(payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
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

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeDetail ...
func EmployeeDetail(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Get employee by ID
	result, err := service.EmployeeDetail(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeList ...
func EmployeeList(c echo.Context) error {

	//Get employee list
	result, err := service.EmployeeList()

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeListByActiveState ...
func EmployeeListByActiveState(c echo.Context) error {
	var (
		active = c.Param("active")
	)

	//Get employee list
	result, err := service.EmployeeListByActiveState(active)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeUpdate ...
func EmployeeUpdate(c echo.Context) error {
	var (
		id      = c.Param("id")
		payload = c.Get("body").(model.EmployeeUpdatePayload)
	)

	//Update employee
	result, err := service.EmployeeUpdate(id, payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeChangeActive ...
func EmployeeChangeActive(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Change employee active sate
	result, err := service.EmployeeChangeActive(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeDelete ...
func EmployeeDelete(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Delete employee by ID
	err := service.EmployeeDelete(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
