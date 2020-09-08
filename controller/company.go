package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CompanyCreate ...
func CompanyCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.CompanyCreatePayload)
	)

	//Create company
	result, err := service.CompanyCreate(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyDetail ...
func CompanyDetail(c echo.Context) error {
	var (
		company = c.Get("company").(model.Company)
	)

	//Get company by ID
	result, err := service.CompanyDetail(company.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyList ...
func CompanyList(c echo.Context) error {

	//Get company list
	result, err := service.CompanyList()

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyUpdate ...
func CompanyUpdate(c echo.Context) error {
	var (
		company = c.Get("company").(model.Company)
		payload = c.Get("body").(model.CompanyUpdatePayload)
	)

	//Update company
	result, err := service.CompanyUpdate(company.ID, payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyChangeActive ...
func CompanyChangeActive(c echo.Context) error {
	var (
		company = c.Get("company").(model.Company)
	)

	//Change company active state
	result, err := service.CompanyChangeActive(company.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyDelete ...
func CompanyDelete(c echo.Context) error {
	var (
		company = c.Get("company").(model.Company)
	)

	//Delete company by ID
	err := service.CompanyDelete(company.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
