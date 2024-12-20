package main

import (
	"api-golang/database"
	"api-golang/models"
	"encoding/json"
	"net/http"
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

	http.ListenAndServe(":8080", nil)
}
