package main

import (
	"fmt"
	"log"
	"os"

	"tugas-deploy/controller"
	"tugas-deploy/database"
	"tugas-deploy/model"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&model.Bioskop{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Bioskop API")
	})

	// Endpoint CRUD
	r.GET("/bioskop", controller.GetAllBioskop)
	r.GET("/bioskop/:id", controller.GetBioskopByID)
	r.POST("/bioskop", controller.CreateBioskop)
	r.PUT("/bioskop/:id", controller.UpdateBioskop)
	r.DELETE("/bioskop/:id", controller.DeleteBioskop)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Running on port " + port)
	log.Fatal(r.Run(":" + port))
}
