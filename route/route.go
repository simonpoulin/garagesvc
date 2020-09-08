package route

import "github.com/labstack/echo/v4"

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	admin(e)
	user(e)
	common(e)
}
