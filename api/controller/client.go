package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/a-kumar5/auth-hub/api/utils"
	"github.com/rs/zerolog/log"
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
			log.Error().
				Err(err).
				Msg("Failed to decode request body")
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// hash the password
		hashedPassword, err := utils.HashPassword(c.Password)
		if err != nil {
			log.Error().
				Err(err).
				Str("client_id", c.ClientId).
				Msg("Failed to hash password")
			http.Error(w, "Failed to create client", http.StatusInternalServerError)
			return
		}
		c.Password = hashedPassword

		log.Debug().
			Str("client_id", c.ClientId).
			Str("name", c.Name).
			Msg("Creating new client")

		c.CreatedAt = time.Now()

		err = db.QueryRow("INSERT INTO clients (client_name, client_id, client_password, created_at) VALUES ($1, $2, $3, $4) RETURNING id", c.Name, c.ClientId, c.Password, c.CreatedAt).Scan(&c.ID)
		if err != nil {
			log.Error().
				Err(err).
				Str("client_id", c.ClientId).
				Msg("Failed to insert client into database")
			http.Error(w, "Failed to create client", http.StatusInternalServerError)
			return
		}

		log.Info().
			Int("id", c.ID).
			Str("client_id", c.ClientId).
			Msg("Client created successfully")

		w.WriteHeader(http.StatusOK)

		// hiding password from sending as part of request
		c.Password = ""
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":      c.ID,
			"client":  c,
			"message": "Client created successfully",
		})
	}
}

func GetClients(db *sql.DB, secretKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Fetching all clients")
		token := r.Header.Get("Authorization")

		// verify token validity
		_, err := utils.VerifyToken(token, secretKey)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Not able to validate the token or token not provided.")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		rows, err := db.Query("SELECT id, client_name, client_id, created_at FROM Clients")
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to query clients from database")
			http.Error(w, "Resource not found", http.StatusNotFound)
			return
		}
		defer rows.Close()

		clients := []Client{}
		for rows.Next() {
			var c Client
			if err := rows.Scan(&c.ID, &c.Name, &c.ClientId, &c.CreatedAt); err != nil {
				log.Error().
					Err(err).
					Msg("Failed to scan client row")
				http.Error(w, "Error retrieving clients", http.StatusInternalServerError)
				return
			}
			clients = append(clients, c)
		}

		log.Info().
			Int("count", len(clients)).
			Msg("Successfully fetched clients")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"client":  clients,
			"message": "Client fetched successfully",
		})
	}
}
