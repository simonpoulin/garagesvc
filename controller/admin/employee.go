package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// EmployeeDetail godoc
//
// @Summary Admin API - Employee detail
// @Description Return employee's details
// @Tags Admin|Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee's ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /admin/employees/{id} [get]
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
// @Summary Admin API - List employees
// @Description Returns a list of all employees
// @Tags Admin|Employees
// @Accept  json
// @Produce  json
// @Param active query string false "Active state"
// @Param name query string false "Employee(s)'s name"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /admin/employees/ [get]
func EmployeeList(c echo.Context) error {
	var (
		active = c.QueryParam("active")
		name   = c.QueryParam("name")
		page   = c.Get("page").(int)
	)

	//Get employee list
	result, err := service.EmployeeList(active, name, page)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// EmployeeUpdate godoc
//
// @Summary Admin API - Employee update
// @Description Update employee's details
// @Tags Admin|Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /admin/employees/{id} [patch]
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

// EmployeeDelete godoc
//
// @Summary Admin API - Employee delete
// @Description Delete an employee
// @Tags Admin|Employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} util.Response
// @Failure 404 {object} util.Response
// @Router /admin/employees/{id} [delete]
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
