package user

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

// CompanyDetail godoc
//
// @Summary User API - Company detail
// @Description Return details of a company
// @Tags User - Companies
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
// @Router /user/companies/{id} [get]
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
// @Summary User API - Company list
// @Description Return a list of companies
// @Tags User - Companies
//
// @Accept  json
// @Produce  json
//
// @Param name query string false "Name keyword"
// @Param active query string false "Active state"
// @Param page query int false "Page number"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /user/companies/ [get]
func CompanyList(c echo.Context) error {
	var (
		name   = c.QueryParam("name")
		page   = c.Get("page").(int)
		active = c.QueryParam("active")
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
