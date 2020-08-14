package validator

import (
	"garagesvc/model"
	"garagesvc/util"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// CustomerValid ...
func CustomerValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CustomerPayload
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

// CustomerLogin ...
func CustomerLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CustomerLoginPayload
		)

		//Bind and parse to struct
		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Validate struct
		_, err := govalidator.ValidateStruct(payload)
		if err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}