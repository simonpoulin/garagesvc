package user

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceDetail godoc
//
// @Summary User API - Service detail
// @Description Return service's details
// @Tags User - Services
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
// @Router /user/services/{id} [get]
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
// @Summary User API - List services
// @Description Returns a list of services
// @Tags User - Services
//
// @Accept  json
// @Produce  json
//
// @Param name query string false "Name keyword"
// @Param companyid query string false "Company's ID"
// @Param active query string false "Active state"
//
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 401 {object} util.Response
// @Failure 404 {object} util.Response
//
// @Security BearerToken
// @Router /user/services/ [get]
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
