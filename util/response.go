package util

import (
	"net/http"

	"github.com/labstack/echo"
)

type response struct {
	Message string
	Data    interface{}
}

// Response200 ...
func Response200(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Thành công!"
	}

	return c.JSON(http.StatusOK, response{
		Message: message,
		Data:    data,
	})
}

// Response400 ...
func Response400(c echo.Context, message string) error {
	if message == "" {
		message = "Không hợp lệ!"
	}

	return c.JSON(http.StatusOK, response{
		Message: message,
	})
}

// Response401 ...
func Response401(c echo.Context, message string) error {
	if message == "" {
		message = "Bạn không có quyền truy cập!"
	}

	return c.JSON(http.StatusUnauthorized, response{
		Message: message,
	})
}
