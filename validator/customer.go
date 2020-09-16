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
		_, err := govalidator.ValidateStruct(payload)

		//Validate struct
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Check phone number is existed

		//Set filter
		filter := bson.M{"phone": payload.Phone}

		//Looking for customer from database
		_, err = dao.CustomerFindOne(filter)
		if err != nil {
			if err.Error() != "mongo: no documents in result" {
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
		_, err := govalidator.ValidateStruct(payload)

		//Validate struct
		if err != nil {
			return util.Response400(c, err.Error())
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
			id       = c.Param("id")
			_id      primitive.ObjectID
			customer model.Customer
		)

		//Bind ID
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set filter
		filter := bson.M{"_id": _id}

		//Validate customer
		customer, err = dao.CustomerFindOne(filter)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("customer", customer)
		return next(c)
	}
}

// CustomerOwner ...
func CustomerOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id           = c.Param("id")
			authcustomer = c.Get("authcustomer").(model.Customer)
			_id          primitive.ObjectID
			customer     model.Customer
		)

		//Bind ID
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set filter
		filter := bson.M{"_id": _id}

		//Validate customer
		customer, err = dao.CustomerFindOne(filter)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Check if a user is also an owner
		if customer.ID != authcustomer.ID {
			err := errors.New("you are not the owner")
			return util.Response401(c, err.Error())
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
			page = c.QueryParam("page")
			p    = 0
			err  error
		)

		//Check valid page param
		if page != "" {
			p, err = strconv.Atoi(page)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}
		c.Set("page", p)

		//Move to next process
		return next(c)
	}
}
