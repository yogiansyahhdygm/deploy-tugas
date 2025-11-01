package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	database "tugas-deploy/config"
	"tugas-deploy/models"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Bioskop{})

	http.HandleFunc("/bioskop", handleBioskop)
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Selamat datang di API Bioskop!")
		})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleBioskop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		var bioskop []models.Bioskop
		database.DB.Find(&bioskop)
		json.NewEncoder(w).Encode(bioskop)

	case "POST":
		var b models.Bioskop
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

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
