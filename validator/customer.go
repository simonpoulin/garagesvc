package validator

import (
	"errors"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerRegister ...
func CustomerRegister(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CustomerRegisterPayload
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
		_, err = dao.CustomerFindOne(filter)
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

// CustomerUpdate ...
func CustomerUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CustomerUpdatePayload
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

// CustomerLogin ...
func CustomerLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CustomerLoginPayload
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

// CustomerCheckExistance ...
func CustomerCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		//Validate customer
		customer, err := CustomerValidate(id)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("customer", customer)
		return next(c)
	}
}

// CustomerFindRequest ...
func CustomerFindRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query model.CustomerQuery
		)

		//Bind and parse to struct
		if err := c.Bind(&query); err != nil {
			return util.Response400(c, err.Error())
		}

		//Set body and move to next process
		c.Set("query", query)
		return next(c)
	}
}

// CustomerValidate ...
func CustomerValidate(customerID string) (customer model.Customer, err error) {

	//Check valid customer ID
	ctmID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return
	}

	//Set filter
	filter := bson.M{"_id": ctmID}

	//Validate customer
	customer, err = dao.CustomerFindOne(filter)

	return
}
