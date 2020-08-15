package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateToken used to create JWT
func CreateToken(userid uint32) (string, error) {
	var err error
	//Creating Access Token
	fmt.Println("tried creating token for id " + string(userid))
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 23).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("jwtsecret")))
	if err != nil {
		return "", err
	}
	return token, nil
}
