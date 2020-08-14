package handlers

import (
	"net/http"

	"github.com/BRO3886/gin-learn/api/middleware"

	"github.com/BRO3886/gin-learn/pkg/user"
	"github.com/gin-gonic/gin"
)

//RegisterUser used to reg user
func RegisterUser(svc user.Service) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		user := &user.User{}
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		user, err := svc.Register(user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		token, err := middleware.CreateToken(user.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		user.Password = ""
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "user created",
			"token":   token,
			"user":    *user,
		})
	}
}

//LoginUser is used gor loggin in user
func LoginUser(svc user.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}
