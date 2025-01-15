package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Client struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ClientId  string    `json:"client_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateClient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c Client
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("Client ID, Name, and Password: %v %v %v", c.ClientId, c.Name, c.Password)

		c.CreatedAt = time.Now()

		err = db.QueryRow("INSERT INTO clients (client_name, client_id, client_password, created_at) VALUES ($1, $2, $3, $4) RETURNING id", c.Name, c.ClientId, c.Password, c.CreatedAt).Scan(&c.ID)
		if err != nil {
			log.Printf("Error inserting into database: %v", err)
			http.Error(w, "Failed to create client", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":      c.ID,
			"client":  c,
			"message": "Client created successfully",
		})
	}
}
