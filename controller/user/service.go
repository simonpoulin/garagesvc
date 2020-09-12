package user

import (
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
