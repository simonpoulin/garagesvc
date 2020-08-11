package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(emId string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "randomstring") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["em_id"] = emId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
