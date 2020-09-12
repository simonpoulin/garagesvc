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

// BookingCreate ...
func BookingCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.BookingCreatePayload
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

// BookingUpdate ...
func BookingUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.BookingUpdatePayload
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

// BookingCheckExistance ...
func BookingCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			_id     primitive.ObjectID
			booking model.Booking
		)

		//Bind ID
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set filter
		filter := bson.M{"_id": _id}

		//Validate booking
		booking, err = dao.BookingFindOne(filter)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("booking", booking)
		return next(c)
	}
}

// BookingFindRequest ...
func BookingFindRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			customerID                    = c.QueryParam("customer_id")
			serviceID                     = c.QueryParam("service_id")
			page                          = c.QueryParam("page")
			status                        = c.QueryParam("status")
			p                             = 0
			ctmID      primitive.ObjectID = [12]byte{}
			svcID      primitive.ObjectID = [12]byte{}
			err        error
		)

		//Check valid page param
		if page != "" {
			p, err = strconv.Atoi(page)
			if err != nil || p < 1 {
				return util.Response400(c, err.Error())
			}
		}
		c.Set("page", p)

		//Check valid page param
		if status != "" {
			_, err = strconv.ParseBool(status)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}

		//Check valid customerID and set body
		if customerID != "" {
			ctmID, err = primitive.ObjectIDFromHex(customerID)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}
		c.Set("customerID", ctmID)

		//Check valid customerID and set body
		if serviceID != "" {
			svcID, err = primitive.ObjectIDFromHex(serviceID)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}
		c.Set("serviceID", svcID)

		//Move to next process
		return next(c)
	}
}

// BookingOwner ...
func BookingOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			booking  = c.Get("booking").(model.Booking)
			customer = c.Get("authcustomer").(model.Customer)
		)

		//Check if a user is also an owner
		if customer.ID != booking.CustomerID {
			err := errors.New("you are not the owner")
			return util.Response401(c, err.Error())
		}

		//Move to next process
		return next(c)
	}
}
