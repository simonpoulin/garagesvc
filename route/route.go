package route

import (
	"garagesvc/route/admin"
	"garagesvc/route/user"

	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	admin.RouteAdmin(e)
	user.RouteUser(e)
	common(e)
}
