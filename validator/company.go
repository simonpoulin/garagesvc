package validator

import (
	"errors"
	"garagesvc/model"
	"garagesvc/service"
	"garagesvc/util"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// CompanyCreate ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CompanyCreatePayload
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

// CompanyUpdate ...
func CompanyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CompanyUpdatePayload
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

// CompanyCheck ...
func CompanyCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			companyID = c.Param("companyID")
		)

		//Validate company ID
		company, err := service.CompanyDetail(companyID)
		if err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Check company status
		if !company.Active {
			err = errors.New("company not active")
			return util.Response400(c, err.Error(), nil)
		}

		return next(c)
	}
}
