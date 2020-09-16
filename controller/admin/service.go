package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceCreate godoc
//
// @Summary Admin API - Service create
// @Description Create a service
// @Tags Admin - Services
//
// @Accept  json
// @Produce  json
//
// @Param ServiceCreatePayload body model.ServiceCreatePayload true "Service Create Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/services/ [post]
func ServiceCreate(c echo.Context) error {
	var (
		payload = c.Get("body").(model.ServiceCreatePayload)
	)

	//Create service
	result, err := service.ServiceCreate(payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceDetail godoc
//
// @Summary Admin API - Service detail
// @Description Return service's details
// @Tags Admin - Services
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Service's ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/services/{id} [get]
func ServiceDetail(c echo.Context) error {
	var (
		svc = c.Get("service").(model.Service)
	)

	//Get service by ID
	result, err := service.ServiceDetail(svc.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceList godoc
//
// @Summary Admin API - List services
// @Description Returns a list of services
// @Tags Admin - Services
//
// @Accept  json
// @Produce  json
//
// @Param name query string false "Name keyword"
// @Param companyid query string false "Company's ID"
// @Param active query string false "Active state"
// @Param page query int false "Page number"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/services/ [get]
func ServiceList(c echo.Context) error {
	var (
		active    = c.QueryParam("active")
		name      = c.QueryParam("name")
		companyID = c.Get("companyID").(primitive.ObjectID)
		page      = c.Get("page").(int)
	)

	//Get service list
	result, err := service.ServiceList(active, name, companyID, page)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceUpdate godoc
//
// @Summary Admin API - Service update
// @Description Update service's details
// @Tags Admin - Services
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Service's ID"
// @Param active query string false "Active state"
// @Param ServiceUpdatePayload body model.ServiceUpdatePayload false "Service Update Payload"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/services/{id} [patch]
func ServiceUpdate(c echo.Context) error {
	var (
		svc     = c.Get("service").(model.Service)
		payload = c.Get("body").(model.ServiceUpdatePayload)
		active  = c.QueryParam("active")
	)

	//Update service
	result, err := service.ServiceUpdate(svc.ID, payload, active)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceDelete godoc
//
// @Summary Admin API - Service delete
// @Description Delete an service
// @Tags Admin - Services
//
// @Accept  json
// @Produce  json
//
// @Param id path string true "Service's ID"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /admin/services/{id} [delete]
func ServiceDelete(c echo.Context) error {
	var (
		svc = c.Get("service").(model.Service)
	)

	//Delete service by ID
	err := service.ServiceDelete(svc.ID)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", nil)
}
