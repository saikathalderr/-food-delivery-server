package main

import (
	"foodDeliveryAppServer/controllers"
	"foodDeliveryAppServer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	// Route setup
	r := gin.Default()
	r.Use(database.DbMiddleware)

	api := r.Group("/api")
	// Health check
	api.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Business
	business := api.Group("/business")

	// Business - Auth
	businessAuth := business.Group("/auth")
	businessAuth.POST("/login", controllers.LoginHandler)
	businessAuth.POST("/sign-up", controllers.SignUpHandler)

	return r
}

func main() {
	// setup routes
	r := setupRouter()

	// Server in 0.0.0.0:3030
	r.Run(":3030")
}
