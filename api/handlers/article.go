package handlers

import (
	"net/http"
	"os"

	"github.com/BRO3886/gin-learn/api"

	"github.com/BRO3886/gin-learn/api/middleware"
	"github.com/BRO3886/gin-learn/pkg/article"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//CreateNewArticle is used to create new articles
func CreateNewArticle(svc article.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		tk := &middleware.Token{}
		_, _ = jwt.ParseWithClaims(tokenHeader, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtsecret")), nil
		})
		article := &article.Article{}

		if err := ctx.ShouldBindJSON(&article); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if article.Title == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": api.ErrNoTitle.Error()})
			ctx.Abort()
			return
		}

		article.UserID = tk.ID

		article, err := svc.CreateArticle(article)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "article created",
			"article": *article,
		})

	}
}

//GetArticlesByUser to return a list of articles created by user
func GetArticlesByUser(svc article.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		tk := &middleware.Token{}
		_, _ = jwt.ParseWithClaims(tokenHeader, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtsecret")), nil
		})

		articles, err := svc.GetUserArticles(uint32(tk.ID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "articles found",
			"articles": *articles,
		})
	}
}
