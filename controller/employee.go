package controller

import (
	"context"
	"fmt"
	model "garagesvc/models"
	util "garagesvc/utils"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func Connect() *mongo.Collection {
	client = util.ConnectDB()
	return client.Database("garagesvc").Collection("employee")
}

func Disconnect(c context.Context) {
	client.Disconnect(c)
}

func EmployeeLogin(c echo.Context) error {
	collection, ctxt := Connect(), context.Background()
	defer Disconnect(ctxt)
	var e model.EmployeeLoginPayload
	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	e.Password = util.Hash(e.Password)
	var em model.Employee
	filter := bson.M{"phone": e.Phone}
	err := collection.FindOne(ctxt, filter).Decode(&em)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if e.Password != em.Password {
		fmt.Print(e.Password + " " + em.Password)
		return c.JSON(http.StatusNotFound, "")
	}
	token, err := util.CreateToken(em.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, token)
}

func EmployeeCreate(c echo.Context) error {
	collection, ctxt := Connect(), context.Background()
	defer Disconnect(ctxt)
	var e model.Employee
	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	e.ID = primitive.NewObjectID()
	e.Active = true
	e.Password = util.Hash(e.Password)
	_, err := collection.InsertOne(ctxt, e)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, e)
}
