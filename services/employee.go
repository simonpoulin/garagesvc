package services

import (
	"context"
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

func EmployeeCreate(payload model.EmployeeCreatePayload, c echo.Context) error {
	collection, ctxt := Connect(), context.Background()
	defer Disconnect(ctxt)
	var e model.Employee
	e.ID = primitive.NewObjectID()
	e.Active = true
	e.Password = util.Hash(payload.Password)
	e.Name = payload.Name
	e.Phone = payload.Phone
	_, err := collection.InsertOne(ctxt, e)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, e)
}

func EmployeeLogin(payload model.EmployeeLoginPayload, c echo.Context) error {
	collection, ctxt := Connect(), context.Background()
	defer Disconnect(ctxt)
	payload.Password = util.Hash(payload.Password)
	var e model.Employee
	filter := bson.M{"phone": payload.Phone}
	err := collection.FindOne(ctxt, filter).Decode(&e)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if payload.Password != e.Password {
		return c.JSON(http.StatusNotFound, "")
	}
	token, err := util.CreateToken(e.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, token)
}
