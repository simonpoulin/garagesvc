package common

import (
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// TestString ...
func TestString(c echo.Context) error {
	var (
		query     model.AppQuery
		findQuery = query.GenerateFindQuery()
		update    bson.M
		filter    bson.M
	)

	c1, err := dao.EmployeeFind(findQuery)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	for i := range c1 {
		//Set filter and data
		cpn := c1[i]
		filter = bson.M{"_id": cpn.ID}
		data := cpn.ConvertToUpdateSearchStringBSON()
		update = bson.M{"$set": data}

		//Update company
		err = dao.EmployeeUpdateOne(filter, update)
		if err != nil {
			return util.Response404(c, err.Error())
		}
	}

	c2, err := dao.CompanyFind(findQuery)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	for i := range c2 {
		//Set filter and data
		cpn := c2[i]
		filter = bson.M{"_id": cpn.ID}
		data := cpn.ConvertToUpdateSearchStringBSON()
		update = bson.M{"$set": data}

		//Update company
		err = dao.CompanyUpdateOne(filter, update)
		if err != nil {
			return util.Response404(c, err.Error())
		}
	}

	c3, err := dao.CustomerFind(findQuery)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	for i := range c3 {
		//Set filter and data
		cpn := c3[i]
		filter = bson.M{"_id": cpn.ID}
		data := cpn.ConvertToUpdateSearchStringBSON()
		update = bson.M{"$set": data}

		//Update company
		err = dao.CustomerUpdateOne(filter, update)
		if err != nil {
			return util.Response404(c, err.Error())
		}
	}

	c4, err := dao.ServiceFind(findQuery)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	for i := range c4 {
		//Set filter and data
		cpn := c4[i]
		filter = bson.M{"_id": cpn.ID}
		data := cpn.ConvertToUpdateSearchStringBSON()
		update = bson.M{"$set": data}

		//Update company
		err = dao.ServiceUpdateOne(filter, update)
		if err != nil {
			return util.Response404(c, err.Error())
		}
	}

	//Return 200
	return util.Response200(c, "", nil)
}

// // TestString ...
// func TestString(c echo.Context) error {

// }
