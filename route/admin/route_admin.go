package admin

import (
	"garagesvc/auth"

	"github.com/labstack/echo/v4"
)

// RouteAdmin ...
func RouteAdmin(e *echo.Echo) {
	group := e.Group("/admin")

	group.Use(auth.LoggedInAsEmployee)

	adminBooking(group)
	adminCompany(group)
	adminCustomer(group)
	adminEmployee(group)
	adminService(group)
}
