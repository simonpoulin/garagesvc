package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// EmployeeCreate godoc
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

// EmployeeDetail godoc
//
// @Summary Employee detail
// @Description Return employee's details
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /employees/{id} [get]
func EmployeeDetail(c echo.Context) error {
	var (
		employee = c.Get("employee").(model.Employee)
	)

	//Get employee by ID
	result, err := service.EmployeeDetail(employee.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeList godoc
//
// @Summary List employees
// @Description Returns list of all employees
// @Tags Employees
// @Accept  json
// @Produce  json
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /employees/ [get]
func EmployeeList(c echo.Context) error {

	//Get employee list
	result, err := service.EmployeeList()

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeListByActiveState godoc
//
// @Summary Employee list by active state
// @Description Return list of all employees by chosen status
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param active path string true "Active state"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /employees/active/{active} [get]
func EmployeeListByActiveState(c echo.Context) error {
	var (
		active = c.Param("active")
	)

	//Get employee list
	result, err := service.EmployeeListByActiveState(active)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeUpdate godoc
//
// @Summary Employee update
// @Description Update employee's details
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /employees/{id} [patch]
func EmployeeUpdate(c echo.Context) error {
	var (
		employee = c.Get("employee").(model.Employee)
		payload  = c.Get("body").(model.EmployeeUpdatePayload)
	)

	//Update employee
	result, err := service.EmployeeUpdate(employee.ID, payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeChangeActive godoc
//
// @Summary Employee update active state
// @Description Update employee's active state
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /employees/{id}/active [patch]
func EmployeeChangeActive(c echo.Context) error {
	var (
		employee = c.Get("employee").(model.Employee)
	)

	//Change employee active sate
	result, err := service.EmployeeChangeActive(employee.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeDelete godoc
//
// @Summary Employee delete
// @Description Delete an employee
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /employees/{id} [delete]
func EmployeeDelete(c echo.Context) error {
	var (
		employee = c.Get("employee").(model.Employee)
	)

	//Delete employee by ID
	err := service.EmployeeDelete(employee.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
