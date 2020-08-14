package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/gin-learn/api/handlers"
	"github.com/BRO3886/gin-learn/pkg/user"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
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
	fmt.Println("Stared backend server")

	db, err := connectDB()
	if err != nil {
		log.Fatal("error conneting to DB")
	}
	db.LogMode(true)
	db.AutoMigrate(&user.User{})
	log.Println("connected to db")

	defer db.Close()

	userRepo := user.NewDatabaseRepo(db)
	userSvc := user.NewService(userRepo)

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		usrGroup := v1.Group("/user")
		{
			usrGroup.POST("/register", handlers.RegisterUser(userSvc))
		}
		// v1.Group("article")
	}

	r.Run(":8000")

}
