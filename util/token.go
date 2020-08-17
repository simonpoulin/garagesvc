package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// UserClaim ...
type UserClaim struct {
	ID string `json:"_id"`
	jwt.StandardClaims
}

// TokenEncode ...
func TokenEncode(ID string) (tokenString string, err error) {

	// Load dotenv for signing key
	err = godotenv.Load()
	if err != nil {
		return "", err
	}

	//Generate token
	claims := &UserClaim{
		ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return
}
