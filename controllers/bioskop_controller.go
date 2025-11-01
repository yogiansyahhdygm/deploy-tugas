package controllers

import (
	"net/http"
	"strconv"

	"tugas-deploy/config"
	"tugas-deploy/models"

	"github.com/gin-gonic/gin"
)

// --- CREATE ---
func CreateBioskop(c *gin.Context) {
	var input models.Bioskop

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan", "data": input})
}

// --- READ ALL ---
func GetAllBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	config.DB.Find(&bioskops)
	c.JSON(http.StatusOK, bioskops)
}

// --- READ ONE ---
func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Bioskop

	if err := config.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// --- UPDATE ---
func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var b models.Bioskop

	if err := config.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	b.Nama = input.Nama
	b.Lokasi = input.Lokasi
	b.Rating = input.Rating
	config.DB.Save(&b)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui", "data": b})
}

// --- DELETE ---
func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	var b models.Bioskop
	if err := config.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	config.DB.Delete(&b)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
