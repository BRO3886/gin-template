package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/gin-learn/api/handlers"
	"github.com/BRO3886/gin-learn/pkg/article"
	"github.com/BRO3886/gin-learn/pkg/user"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func connectDB() (*gorm.DB, error) {

	//Heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	return gorm.Open("sqlite3", "test.db")
}

func main() {
	log.Println("Stared backend server")
	gin.SetMode(gin.ReleaseMode)

	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		log.Fatal("error conneting to DB")
	}
	if os.Getenv("LOG_MODE") == "FALSE" {
		db.LogMode(false)
	}
	db.AutoMigrate(&user.User{}, &article.Article{})
	log.Println("connected to db")

	defer db.Close()

	userRepo := user.NewDatabaseRepo(db)
	userSvc := user.NewService(userRepo)
	articleRepo := article.NewDatabaseRepo(db)
	articleSvc := article.NewService(articleRepo)

	r := gin.Default()

	r.HandleMethodNotAllowed = true
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		handlers.MakeUserHandler(v1, userSvc)
		handlers.MakeArticleHandlers(v1, articleSvc)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}

	log.Println("runnning on port " + port)

	r.Run(":" + port)

}
