package controller

import (
	model "garagesvc/models"
	service "garagesvc/services"

	"github.com/labstack/echo"
)

func EmployeeCreate(c echo.Context) error {
	result := service.EmployeeCreate(c.Get("body").(model.EmployeeCreatePayload), c)
	return result
}

func EmployeeLogin(c echo.Context) error {
	result := service.EmployeeLogin(c.Get("body").(model.EmployeeLoginPayload), c)
	return result
}
