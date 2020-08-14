package authentication

import (
	"errors"
	"garagesvc/service"
	"garagesvc/util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Authenticate ...
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//Decode token and bind data
		var (
			token     = c.Get("token").(*jwt.Token)
			userClaim = token.Claims.(*util.UserClaim)
		)

		//Verify employee by ID
		_, err := service.EmployeeFindOne(userClaim.ID)
		if err != nil {
			return util.Response400(c, err.Error(), nil)
		}

		//Verify expired time
		if userClaim.VerifyExpiresAt(time.Now().Unix(), true) {
			err = errors.New("token expired")
			return util.Response400(c, err.Error(), nil)
		}

		//Set body and move to next process
		return next(c)
	}
}
