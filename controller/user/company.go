package user

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
)

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
	var (
		name = c.QueryParam("name")
		page = c.Get("page").(int)
	)

	//Get company list
	result, err := service.CompanyList(name, page)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}