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
			id = c.Param("id")
		)

		//Validate company
		company, err := CompanyValidate(id)
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
			query model.CompanyQuery
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

// CompanyValidate ...
func CompanyValidate(companyID string) (company model.Company, err error) {

	//Check valid company ID
	cpnID, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		return
	}

	//Set filter
	filter := bson.M{"_id": cpnID}

	//Validate company
	company, err = dao.CompanyFindOne(filter)

	return
}
