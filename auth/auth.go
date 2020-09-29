package auth

import (
	"errors"
	"garagesvc/config"
	"garagesvc/util"
	"garagesvc/validator"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// LoggedInAsCustomer ...
func LoggedInAsCustomer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Get token string from header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return util.Response401(c, "")
		}

		authToken := strings.Split(authHeader, " ")
		if len(authToken) < 2 {
			return util.Response401(c, "")
		}

		tokenString := authToken[1]
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
			env := config.GetENV()
			key = []byte(env.CustomerKey)
			return
		})
		if err != nil {
			return util.Response401(c, err.Error())
		}
		userClaim, ok := token.Claims.(*util.UserClaim)
		if !ok || !token.Valid {
			return util.Response401(c, "Invalid token")
		}

		//Validate customer
		customer, err := validator.CustomerValidate(userClaim.ID)
		if err != nil {
			return util.Response401(c, err.Error())
		}

		//Set body and move to next process
		c.Set("authcustomer", customer)
		return next(c)
	}
}

// LoggedInAsEmployee ...
func LoggedInAsEmployee(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Get token string from header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return util.Response401(c, "")
		}

		authToken := strings.Split(authHeader, " ")
		if len(authToken) < 2 {
			return util.Response401(c, "")
		}

		tokenString := authToken[1]
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
			env := config.GetENV()
			key = []byte(env.EmployeeKey)
			return
		})
		if err != nil {
			return util.Response401(c, err.Error())
		}
		userClaim, ok := token.Claims.(*util.UserClaim)
		if !ok || !token.Valid {
			return util.Response401(c, "Invalid token")
		}

		//Validate employee
		employee, err := validator.EmployeeValidate(userClaim.ID)
		if err != nil {
			return util.Response401(c, err.Error())
		}
		//Set body and move to next process
		c.Set("authemployee", employee)
		return next(c)
	}
}
