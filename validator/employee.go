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

// EmployeeRegister ...
func EmployeeRegister(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.EmployeeRegisterPayload
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

		//Check phone number is existed
		//Set filter
		filter := bson.M{"phone": payload.Phone}

		//Looking for customer from database
		_, err = dao.EmployeeFindOne(filter)
		if err != nil {
			if !util.IsEmptyListError(err) {
				return util.Response400(c, err.Error())
			}
		} else {
			err = errors.New("phone number existed")
			return util.Response400(c, err.Error())
		}

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// EmployeeLogin ...
func EmployeeLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.EmployeeLoginPayload
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

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// EmployeeUpdate ...
func EmployeeUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.EmployeeUpdatePayload
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

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// EmployeeCheckExistance ...
func EmployeeCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		//Validate employee
		employee, err := EmployeeValidate(id)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("employee", employee)
		return next(c)
	}
}

// EmployeeFindRequest ...
func EmployeeFindRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query model.EmployeeQuery
			err   error
		)

		//Bind and parse to struct
		if err := c.Bind(&query); err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate active param
		if query.Active != "" {
			if query.Active != "all" && query.Active != "active" && query.Active != "inactive" {
				err = errors.New("invalid active query param")
				return util.Response400(c, err.Error())
			}
		}

		//Set body and move to next process
		c.Set("query", query)
		return next(c)
	}
}

// EmployeeValidate ...
func EmployeeValidate(employeeID string) (employee model.Employee, err error) {

	//Check valid employee ID
	empID, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return
	}

	//Set filter
	filter := bson.M{"_id": empID}

	//Validate employee
	employee, err = dao.EmployeeFindOne(filter)

	return
}
