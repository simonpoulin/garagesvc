package auth

import (
	"errors"
	"garagesvc/service"
	"garagesvc/util"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// IsLoggedIn ...
func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Get token string from header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return util.Response401(c, "")
		}

		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			return util.Response401(c, "")
		}

		//Get data and parse claim
		token, err := jwt.ParseWithClaims(tokenString, &util.UserClaim{}, func(token *jwt.Token) (key interface{}, err error) {
			// Validate alg
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err = errors.New("Unexpected signing method: " + token.Header["alg"].(string))
				return
			}
			err = godotenv.Load()
			key = []byte(os.Getenv("ACCESS_SECRET"))
			return
		})
		if err != nil {
			return util.Response401(c, err.Error())
		}
		userClaim, ok := token.Claims.(*util.UserClaim)
		if !ok || !token.Valid {
			return util.Response401(c, "Invalid token")
		}

		//Verify employee by ID
		_, err = service.EmployeeDetail(userClaim.ID)
		if err != nil {
			return util.Response401(c, err.Error())
		}

		//Set body and move to next process
		c.Set("body", userClaim.ID)
		return next(c)
	}
}
