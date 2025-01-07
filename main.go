package main

import (
	"api-golang/database"
	"api-golang/models"
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	db := database.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var countries []models.Country
		results := make([][]models.Country, 0)

		for i := 0; i < 100; i++ {
			db.Find(&countries)
			results = append(results, countries)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
