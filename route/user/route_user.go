package user

import (
	"garagesvc/auth"

	"github.com/labstack/echo/v4"
)

// RouteUser ...
func RouteUser(e *echo.Echo) {
	group := e.Group("/user")

	group.Use(auth.LoggedInAsCustomer)

	userBooking(group)
	userCompany(group)
	userCustomer(group)
	userService(group)
}
