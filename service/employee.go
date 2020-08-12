package service

import (
	"context"
	"errors"
	"garagesvc/model"
	"garagesvc/module/mongodb"
	"garagesvc/util"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeCreate ...
func EmployeeCreate(payload model.EmployeeCreatePayload, c echo.Context) (e model.Employee, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	e.ID = primitive.NewObjectID()
	e.Active = true
	e.Password = util.Hash(payload.Password)
	e.Name = payload.Name
	e.Phone = payload.Phone

	//Insert to database
	_, err = employeeCol.InsertOne(ctxt, e)
	return
}

// EmployeeLogin ...
func EmployeeLogin(payload model.EmployeeLoginPayload, c echo.Context) (token string, err error) {
	var (
		employeeCol = mongodb.EmployeeCol()
		ctxt        = context.Background()
	)

	//Check phone number existance
	payload.Password = util.Hash(payload.Password)
	var e model.Employee
	filter := bson.M{"phone": payload.Phone}
	err = employeeCol.FindOne(ctxt, filter).Decode(&e)
	if err != nil {
		return "", err
	}

	//Check password match
	if payload.Password != e.Password {
		return "", errors.New("password not match")
	}
	token, err = util.TokenEncode(e.ID.Hex())
	return
}
