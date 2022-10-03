package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber int
	Address     string
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")
	// Health check
	api.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return r
}

func connectDb() {
	db, err := gorm.Open(postgres.Open("postgres://saikathalder:saikat123@localhost:5432/food-delivery"), &gorm.Config{})
	if err != nil {
		panic("Faild to connect database!")
	}

	db.AutoMigrate(&User{})
}

func main() {
	// setup connection with db
	connectDb()

	// setup routes
	r := setupRouter()

	// Server in 0.0.0.0:3030
	r.Run(":3030")
}
