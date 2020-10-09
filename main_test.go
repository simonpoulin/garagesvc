package main

import (
	"garagesvc/config"
	"garagesvc/module/mongodb"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Init()
	mongodb.Connect()
}

var (
	resJSON = `{"message": "Thành công!","data": {"_id": "5f7d600d7e3a501817ee214a","name": "Na","phone": "0909121211","password": "fEqNCco3Yq9h5ZUglD3CZJT4lBs=","address": "To 13 - Thon 2 - xa Binh Giang - huyen Thang Binh, , ","resourceid": "000000000000000000000000","smallimage": "5f7691bee77eedd2166d4343.png","mediumimage": "5f7691bee77eedd2166d4342.png","largeimage": "5f7691bee77eedd2166d4341.png"}}`
)

func TestGetCustomer(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customers/:id")
	c.SetParamNames("id")
	c.SetParamValues("5f7d600d7e3a501817ee214a")

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, resJSON, rec.Body.String())
}
