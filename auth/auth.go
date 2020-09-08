package auth

import (
	"errors"
	"garagesvc/config"
	"garagesvc/dao"
	"garagesvc/util"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

		//Set filter
		customerID, err := primitive.ObjectIDFromHex(userClaim.ID)
		if err != nil {
			return util.Response401(c, err.Error())
		}
		filter := bson.M{"_id": customerID}

		//Verify customer by ID
		customer, err := dao.CustomerFindOne(filter)
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
		//Set filter
		employeeID, err := primitive.ObjectIDFromHex(userClaim.ID)
		if err != nil {
			return util.Response401(c, err.Error())
		}
		filter := bson.M{"_id": employeeID}

		//Verify employee by ID
		employee, err := dao.EmployeeFindOne(filter)
		if err != nil {
			return util.Response401(c, err.Error())
		}

		//Set body and move to next process
		c.Set("authemployee", employee)
		return next(c)
	}
}
