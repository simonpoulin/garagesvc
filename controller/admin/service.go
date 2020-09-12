package admin

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceCreate ...
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

// ServiceDetail ...
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

// ServiceList ...
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

// ServiceUpdate ...
func ServiceUpdate(c echo.Context) error {
	var (
		svc     = c.Get("service").(model.Service)
		payload = c.Get("body").(model.ServiceUpdatePayload)
	)

	//Update service
	result, err := service.ServiceUpdate(svc.ID, payload)

	//If error, return 404
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Return 200
	return util.Response200(c, "", result)
}

// ServiceDelete ...
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
