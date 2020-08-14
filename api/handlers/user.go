package handlers

import (
	"net/http"

	"github.com/BRO3886/gin-learn/api/middleware"

	"github.com/BRO3886/gin-learn/pkg/user"
	"github.com/gin-gonic/gin"
)

type usrHandler struct {
	svc user.Service
}

//NewUserHandler handles user routes
func NewUserHandler(svc user.Service) usrHandler {
	return usrHandler{svc: svc}
}

//RegisterUser used to reg user
func (handler usrHandler) RegisterUser(ctx *gin.Context) {

	user := &user.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := handler.svc.Register(user)
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
