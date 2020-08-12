package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// TokenEncode ...
func TokenEncode(emID string) (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["em_id"] = emID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	err = godotenv.Load()
	if err != nil {
		return "", err
	}
	token, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return
}
