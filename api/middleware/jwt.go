package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/BRO3886/gin-learn/api"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Token struct
type Token struct {
	ID uint32 `json:"id"`
	jwt.StandardClaims
}

//BasicJWTAuth auth token checker
func BasicJWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": api.ErrTokenMissing.Error()})
			ctx.Abort()
			return
		}
		tk := &Token{}
		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtsecret")), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusForbidden, gin.H{"message": api.ErrInvalidToken.Error()})
			ctx.Abort()
			return
		}

		// _, err = svc.GetUserByID(tk.ID)
		// if err != nil {
		// 	ctx.JSON(http.StatusForbidden, gin.H{"message": api.ErrInvalidToken.Error()})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()
	}
}

//CreateToken used to create JWT
func CreateToken(userid uint32) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["ID"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 23).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("jwtsecret")))
	if err != nil {
		return "", err
	}
	return token, nil
}
