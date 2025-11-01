package main

import (
	"net/http"
	"os"

	"tugas-deploy/config"
	"tugas-deploy/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	config.ConnectDatabase()

	// Router
	r := gin.Default()

	// Test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running 🚀"})
	})

	// CRUD routes
	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetAllBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopByID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
