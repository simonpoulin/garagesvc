package route

import "github.com/labstack/echo"

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	booking(e)
	company(e)
	employee(e)
	customer(e)
	service(e)
}
