package validator

import (
	"errors"
	"fmt"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

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
			err     error
		)

		//Bind and parse to struct
		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate struct
		_, err = govalidator.ValidateStruct(payload)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate customer
		customer, err := CustomerValidate(payload.CustomerID)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Validate service
		service, err := ServiceValidate(payload.ServiceID)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		payload.CustomerObjectID = customer.ID
		payload.ServiceObjectID = service.ID
		c.Set("body", payload)
		return next(c)
	}
}

// BookingUpdate ...
func BookingUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.BookingUpdatePayload
			err     error
		)

		//Bind and parse to struct
		if err := c.Bind(&payload); err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate struct
		_, err = govalidator.ValidateStruct(payload)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate service
		service, err := ServiceValidate(payload.ServiceID)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		payload.ServiceObjectID = service.ID
		c.Set("body", payload)
		return next(c)
	}
}

// BookingCheckExistance ...
func BookingCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		//Validate booking
		booking, err := BookingValidate(id)
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
			query    model.BookingQuery
			err      error
			service  model.Service
			customer model.Customer
		)

		//Bind and parse to struct
		if err = c.Bind(&query); err != nil {
			return util.Response400(c, err.Error())
		}

		//Validate service
		if query.ServiceID != "" {
			service, err = ServiceValidate(query.ServiceID)
			if err != nil {
				fmt.Println("yeet!!!!!")
				return util.Response400(c, err.Error())
			}
		}

		//Validate customer
		if query.CustomerID != "" {
			customer, err = CustomerValidate(query.CustomerID)
			if err != nil {
				fmt.Println("yeet!!!!!-------")
				return util.Response400(c, err.Error())
			}
		}

		//Set body and move to next process
		query.ServiceObjectID = service.ID
		query.CustomerObjectID = customer.ID
		c.Set("query", query)
		return next(c)
	}
}

// BookingOwner ...
func BookingOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id       = c.Param("id")
			customer = c.Get("authcustomer").(model.Customer)
		)

		//Validate booking
		booking, err := BookingValidate(id)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Check if a user is also an owner
		if customer.ID != booking.CustomerID {
			err := errors.New("you are not the owner")
			return util.Response401(c, err.Error())
		}

		//Set body and move to next process
		c.Set("booking", booking)

		//Move to next process
		return next(c)
	}
}

// BookingValidate ...
func BookingValidate(bookingID string) (booking model.Booking, err error) {

	//Check valid booking ID
	bkID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return
	}

	//Set filter
	filter := bson.M{"_id": bkID}

	//Validate booking
	booking, err = dao.BookingFindOne(filter)

	return
}
