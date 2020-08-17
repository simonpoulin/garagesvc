package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo"
)

// CompanyCreate ...
func CompanyCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CompanyCreatePayload)
	)

	//Create company
	result, err := service.CompanyCreate(payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyDetail ...
func CompanyDetail(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Get company by ID
	result, err := service.CompanyDetail(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyList ...
func CompanyList(c echo.Context) error {

	//Get company list
	result, err := service.CompanyList()

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyUpdate ...
func CompanyUpdate(c echo.Context) error {
	var (
		id      = c.Param("id")
		payload = c.Get("body").(model.CompanyUpdatePayload)
	)

	//Update company
	result, err := service.CompanyUpdate(id, payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyChangeActive ...
func CompanyChangeActive(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Change company active state
	result, err := service.CompanyChangeActive(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyDelete ...
func CompanyDelete(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Delete company by ID
	err := service.CompanyDelete(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
