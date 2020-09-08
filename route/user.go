package route

import (
	"garagesvc/auth"

	"github.com/labstack/echo/v4"
)

func user(e *echo.Echo) {
	group := e.Group("/user")

	group.Use(auth.LoggedInAsCustomer)

	userBooking(group)
	userCompany(group)
	userCustomer(group)
	userService(group)
}
