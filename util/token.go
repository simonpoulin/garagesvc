package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserClaim ...
type UserClaim struct {
	ID string `json:"_id"`
	jwt.StandardClaims
}

// TokenEncode ...
func TokenEncode(ID string, key string) (tokenString string, err error) {

	//Generate token
	claims := &UserClaim{
		ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(key))
	return
}
