package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/gin-learn/api/handlers"
	"github.com/BRO3886/gin-learn/pkg/article"
	"github.com/BRO3886/gin-learn/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func connectDB() (*gorm.DB, error) {

	//Heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	//local
	conn, err := pq.ParseURL(os.Getenv("DB_URI"))
	fmt.Println(conn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return gorm.Open("postgres", conn)
}

func main() {
	if os.Getenv("ON_SERVER") != "True" {
		// Loading the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	log.Println("Started gin backend server")
	gin.SetMode(gin.ReleaseMode)

	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		log.Fatal("error connecting to DB")
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
		port = "8080"
	}

	log.Println("runnning on port " + port)

	r.Run(":" + port)

}
