package controller

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo"
)

// ServiceCreate ...
func ServiceCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.ServiceCreatePayload)
	)

	//Create service
	result, err := service.ServiceCreate(payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceDetail ...
func ServiceDetail(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Get service by ID
	result, err := service.ServiceDetail(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceList ...
func ServiceList(c echo.Context) error {

	//Get service list
	result, err := service.ServiceList()

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceListByCompanyID ...
func ServiceListByCompanyID(c echo.Context) error {
	var (
		companyID = c.Param("companyID")
	)

	//Get service list
	result, err := service.ServiceListByCompanyID(companyID)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceListByActiveState ...
func ServiceListByActiveState(c echo.Context) error {
	var (
		active = c.Param("active")
	)

	//Get service list
	result, err := service.ServiceListByActiveState(active)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceUpdate ...
func ServiceUpdate(c echo.Context) error {
	var (
		id      = c.Param("id")
		payload = c.Get("body").(model.ServiceUpdatePayload)
	)

	//Update service
	result, err := service.ServiceUpdate(id, payload)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceChangeActive ...
func ServiceChangeActive(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Change service active state
	result, err := service.ServiceChangeActive(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceDelete ...
func ServiceDelete(c echo.Context) error {
	var (
		id = c.Param("id")
	)

	//Delete service by ID
	err := service.ServiceDelete(id)

	//If error, return 400
	if err != nil {
		return util.Response400(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
