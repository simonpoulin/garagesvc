package common

import (
	"fmt"
	"garagesvc/dao"
	"garagesvc/model"
	"garagesvc/util"
	"io"
	"os"

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

// TestUpload ...
func TestUpload(c echo.Context) error {
	var (
		path string = "assets/img"
	)

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(file)
		return util.Response404(c, err.Error())
	}
	fmt.Println("1")
	src, err := file.Open()
	if err != nil {
		return util.Response404(c, err.Error())
	}
	defer src.Close()
	fmt.Println("2")
	if _, err := os.Stat(path); err != nil || os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		fmt.Println("Path not exist")
	}
	fmt.Println("3")
	path = path + "/" + file.Filename

	dst, err := os.Create(path)
	if err != nil {
		return util.Response404(c, err.Error())
	}
	fmt.Println("4")
	defer dst.Close()
	fmt.Println("Source: ", src)
	fmt.Println("Destination: ", dst)
	fmt.Println("Destination: ", file.Filename)

	if _, err = io.Copy(dst, src); err != nil {
		return util.Response404(c, err.Error())
	}

	return util.Response200(c, "<p>File %s uploaded successfully.</p>", nil)
}
