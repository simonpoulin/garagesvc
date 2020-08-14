package validator

import (
	"errors"
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// ServiceCreate ...
func ServiceCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.ServiceCreatePayload
		)

		//Bind and parse to struct
		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error(), nil)
		}
		_, err := govalidator.ValidateStruct(payload)

		//Validate struct
		if err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// ServiceUpdate ...
func ServiceUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.ServiceUpdatePayload
		)

		//Bind and parse to struct
		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error(), nil)
		}
		_, err := govalidator.ValidateStruct(payload)

		//Validate struct
		if err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// ServiceCheck ...
func ServiceCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			serviceID = c.Param("serviceID")
		)

		//Validate service ID
		svc, err := service.ServiceDetail(serviceID)
		if err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Check service status
		if !svc.Active {
			err = errors.New("service not active")
			return util.Response400(c, err.Error(), nil)
		}

		return next(c)
	}
}
