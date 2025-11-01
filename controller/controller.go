package controller

import (
	"net/http"

	"tugas-deploy/database"
	"tugas-deploy/model"

	"github.com/gin-gonic/gin"
)

func GetAllBioskop(c *gin.Context) {
	var bioskop []model.Bioskop
	database.DB.Order("id asc").Find(&bioskop)
	c.JSON(http.StatusOK, bioskop)
}

func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var b model.Bioskop
	if err := database.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		return
	}
	c.JSON(http.StatusOK, b)
}

func CreateBioskop(c *gin.Context) {
	var b model.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if b.Nama == "" || b.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}
	database.DB.Create(&b)
	c.JSON(http.StatusOK, b)
}

func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var existing model.Bioskop
	if err := database.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		return
	}

	var updated model.Bioskop
	if updated.Nama == "" || updated.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&existing)
	c.JSON(http.StatusOK, gin.H{"message": "Data bioskop berhasil diperbarui", "data": existing})
}

func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")
	var b model.Bioskop
	if err := database.DB.Delete(&b, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop deleted successfully"})
}
