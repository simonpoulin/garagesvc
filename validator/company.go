package validator

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CompanyCreate ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CompanyCreatePayload
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

// CompanyUpdate ...
func CompanyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.CompanyUpdatePayload
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

// CompanyCheckExistance ...
func CompanyCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			_id     primitive.ObjectID
			company model.Company
		)

		//Bind ID
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set filter
		filter := bson.M{"_id": _id}

		//Validate company
		company, err = dao.CompanyFindOne(filter)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("company", company)
		return next(c)
	}
}

// CompanyFindRequest ...
func CompanyFindRequest(next echo.HandlerFunc) echo.HandlerFunc {
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
