package route

import "github.com/labstack/echo"

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	booking(e)
	company(e)
	employee(e)
	service(e)
	user(e)
}
