package database

import (
	"foodDeliveryAppServer/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://saikathalder:saikat123@localhost:5432/food-delivery"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.LocalBusiness{})
	db.AutoMigrate(&models.MenuItem{})

	return db
}

func DbMiddleware(ctx *gin.Context) {
	// setup connection with db
	db := ConnectDb()

	// If no db then it should throw error
	if db == nil {
		panic("Database instance not found!")
	}

	// Setting db as a variable in gin
	ctx.Set("db", db)
	ctx.Next()
}
