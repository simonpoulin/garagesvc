package validator

import (
	"errors"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceCreate ...
func ServiceCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.ServiceCreatePayload
		)

		//Bind and parse to struct
		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate struct
		_, err := govalidator.ValidateStruct(payload)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate company
		company, err := CompanyValidate(payload.CompanyID)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		payload.CompanyObjectID = company.ID
		c.Set("body", payload)
		return next(c)
	}
}

// ServiceUpdate ...
func ServiceUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.ServiceUpdatePayload
			active  = c.QueryParam("active")
		)

		//If active query param not empty, ignore payload
		if active != "" {
			_, err := strconv.ParseBool(active)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		} else {
			//Bind and parse to struct
			if err := c.Bind(&payload); err != nil {
				return util.Response400(c, err.Error())
			}
			_, err := govalidator.ValidateStruct(payload)

			//Validate struct
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}

		//Validate resource
		if payload.ResourceID != "" {
			resource, err := ResourceValidate(payload.ResourceID)
			if err != nil {
				return util.Response400(c, err.Error())
			}
			payload.ResourceObjectID = resource.ID
		}

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// ServiceCheckExistance ...
func ServiceCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		//Validate service
		service, err := ServiceValidate(id)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("service", service)
		return next(c)
	}
}

// ServiceFindRequest ...
func ServiceFindRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query   model.ServiceQuery
			company model.Company
			err     error
		)

		//Bind and parse to struct
		if err = c.Bind(&query); err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate company
		if query.CompanyID != "" {
			company, err = CompanyValidate(query.CompanyID)
			if err != nil {
				return util.Response404(c, err.Error())
			}
		}

		//Validate active param
		if query.Active != "" {
			if query.Active != "all" && query.Active != "active" && query.Active != "inactive" {
				err = errors.New("invalid active query param")
				return util.Response400(c, err.Error())
			}
		}

		//Set body and move to next process
		query.CompanyObjectID = company.ID
		c.Set("query", query)
		return next(c)
	}
}

// ServiceValidate ...
func ServiceValidate(serviceID string) (service model.Service, err error) {

	//Check valid service ID
	svcID, err := primitive.ObjectIDFromHex(serviceID)
	if err != nil {
		return
	}

	//Set filter
	filter := bson.M{"_id": svcID}

	//Validate service
	service, err = dao.ServiceFindOne(filter)

	return
}
