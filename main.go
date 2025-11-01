package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	database "tugas-deploy/config"
	"tugas-deploy/model"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&model.Bioskop{})

	http.HandleFunc("/bioskop", bioskopHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func bioskopHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		if id != "" {
			getBioskopByID(w, id)
		} else {
			getAllBioskop(w)
		}
	case "POST":
		createBioskop(w, r)
	case "PUT":
		updateBioskop(w, r)
	case "DELETE":
		deleteBioskop(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// --------------------- HANDLER FUNCTIONS ---------------------

func getAllBioskop(w http.ResponseWriter) {
	var bioskop []model.Bioskop
	database.DB.Find(&bioskop)
	json.NewEncoder(w).Encode(bioskop)
}

func getBioskopByID(w http.ResponseWriter, id string) {
	var b model.Bioskop
	if err := database.DB.First(&b, id).Error; err != nil {
		http.Error(w, "Bioskop not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(b)
}

func createBioskop(w http.ResponseWriter, r *http.Request) {
	var b model.Bioskop
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if b.Nama == "" || b.Lokasi == "" {
		http.Error(w, "Nama dan Lokasi tidak boleh kosong", http.StatusBadRequest)
		return
	}
	database.DB.Create(&b)
	json.NewEncoder(w).Encode(b)
}

func updateBioskop(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var existing model.Bioskop
	if err := database.DB.First(&existing, id).Error; err != nil {
		http.Error(w, "Bioskop not found", http.StatusNotFound)
		return
	}

	var updated model.Bioskop
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updated.Nama != "" {
		existing.Nama = updated.Nama
	}
	if updated.Lokasi != "" {
		existing.Lokasi = updated.Lokasi
	}
	if updated.Rating != 0 {
		existing.Rating = updated.Rating
	}

	database.DB.Save(&existing)
	json.NewEncoder(w).Encode(existing)
}

func deleteBioskop(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&model.Bioskop{}, idNum)
	if result.RowsAffected == 0 {
		http.Error(w, "Bioskop not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Bioskop deleted successfully"})
}
