package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CompanyCreate godoc
//
// @Summary Admin API - Company create
// @Description Create a company
// @Tags Admin - Companies
//
// @Accept  json
// @Produce  json
//
// @Param CompanyCreatePayload body model.CompanyCreatePayload true "Company Create Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/companies/ [post]
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

// CompanyDetail godoc
//
// @Summary Admin API - Company detail
// @Description Return details of a company
// @Tags Admin - Companies
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Company's ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/companies/{id} [get]
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

// CompanyList godoc
//
// @Summary Admin API - Company list
// @Description Return a list of companies
// @Tags Admin - Companies
//
// @Accept  json
// @Produce  json
//
// @Param name query string false "Name keyword"
// @Param active query string false "Active state"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/companies/ [get]
func CompanyList(c echo.Context) error {
	var (
		name   = c.QueryParam("name")
		active = c.QueryParam("active")
		page   = c.Get("page").(int)
	)

	//Get company list
	result, err := service.CompanyList(name, page, active)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyUpdate godoc
//
// @Summary Admin API - Company update
// @Description Update company's details
// @Tags Admin - Companies
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Company's ID"
// @Param active query string false "Active state"
// @Param CompanyUpdatePayload body model.CompanyUpdatePayload false "Company Update Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/companies/{id} [patch]
func CompanyUpdate(c echo.Context) error {
	var (
		company = c.Get("company").(model.Company)
		payload = c.Get("body").(model.CompanyUpdatePayload)
		active  = c.QueryParam("active")
	)

	//Update company
	result, err := service.CompanyUpdate(company.ID, payload, active)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// CompanyDelete godoc
//
// @Summary Admin API - Company delete
// @Description Delete a company
// @Tags Admin - Companies
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Company's ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/compnanies/{id} [delete]
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
