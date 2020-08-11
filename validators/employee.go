package validators

import (
	model "garagesvc/models"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

func EmployeeCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.EmployeeCreatePayload
		if err := c.Bind(&payload); err != nil {
			return c.JSON(400, model.Response{
				Message: "Dữ liệu không hợp lệ",
			})
		}
		_, err := govalidator.ValidateStruct(payload)
		if err != nil {
			return c.JSON(400, model.Response{
				Message: err.Error(),
			})
		}
		c.Set("body", payload)
		return next(c)
	}
}

func EmployeeLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.EmployeeLoginPayload
		if err := c.Bind(&payload); err != nil {
			return c.JSON(400, model.Response{
				Message: "Dữ liệu không hợp lệ",
			})
		}
		_, err := govalidator.ValidateStruct(payload)
		if err != nil {
			return c.JSON(400, model.Response{
				Message: err.Error(),
			})
		}
		c.Set("body", payload)
		return next(c)
	}
}
