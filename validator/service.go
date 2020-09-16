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

// ServiceCreate ...
func ServiceCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.ServiceCreatePayload
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

		//Check valid company ID
		cpnID, err := primitive.ObjectIDFromHex(payload.CompanyID)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set filter
		filter := bson.M{"_id": cpnID}

		//Validate company
		_, err = dao.CompanyFindOne(filter)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set IDs for payload
		payload.CompanyObjectID = cpnID

		//Set body and move to next process
		c.Set("body", payload)
		return next(c)
	}
}

// ServiceUpdate ...
func ServiceUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.ServiceUpdatePayload
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

// ServiceCheckExistance ...
func ServiceCheckExistance(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			_id     primitive.ObjectID
			service model.Service
		)

		//Bind ID
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return util.Response400(c, err.Error())
		}

		//Set filter
		filter := bson.M{"_id": _id}

		//Validate service
		service, err = dao.ServiceFindOne(filter)
		if err != nil {
			return util.Response404(c, err.Error())
		}

		//Set body and move to next process
		c.Set("service", service)
		return next(c)
	}
}

// ServiceFindRequest ...
func ServiceFindRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			companyID                    = c.QueryParam("company_id")
			page                         = c.QueryParam("page")
			active                       = c.QueryParam("active")
			p                            = 0
			cpnID     primitive.ObjectID = [12]byte{}
			err       error
		)

		//Check valid page param
		if page != "" {
			p, err = strconv.Atoi(page)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}

		//Check valid active param
		if active != "" {
			_, err = strconv.ParseBool(active)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}

		//Check valid companyID param
		if companyID != "" {
			cpnID, err = primitive.ObjectIDFromHex(companyID)
			if err != nil {
				return util.Response400(c, err.Error())
			}
		}

		//Set body and move to next process
		c.Set("page", p)
		c.Set("companyID", cpnID)
		return next(c)
	}
}
