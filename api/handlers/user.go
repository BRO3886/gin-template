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
			ctx.Abort()
			return
		}
		user, err := svc.Register(user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		token, err := middleware.CreateToken(uint32(user.ID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
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
		user := &user.User{}
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		user, err := svc.Login(user.Email, user.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		token, err := middleware.CreateToken(uint32(user.ID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		user.Password = ""
		ctx.JSON(http.StatusOK, gin.H{
			"message": "login success",
			"token":   token,
			"user":    *user,
		})
		return
	}
}

//GetUserDetails returns user details
func GetUserDetails(svc user.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user, err := svc.GetUserByEmail(ctx.Query("email"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		user.Password = ""
		ctx.JSON(http.StatusFound, gin.H{
			"message": "user found",
			"user":    *user,
		})
		return
	}
}

// MakeUserHandler defines the routes for user
func MakeUserHandler(r *gin.RouterGroup, userSvc user.Service) {
	usrGroup := r.Group("/user")
	{
		usrGroup.POST("/register", RegisterUser(userSvc))
		usrGroup.POST("/login", LoginUser(userSvc))
		usrGroup.Use(middleware.BasicJWTAuth())
		{
			usrGroup.GET("/getdetails", GetUserDetails(userSvc))
		}
	}
}
