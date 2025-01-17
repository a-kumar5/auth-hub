package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/a-kumar5/auth-hub/api/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Token struct {
	ID        uuid.UUID `json:"id"`
	ClientId  string    `json:"client_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func CreateToken(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			ClientId string `json:"client_id"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&body)

		log.Info().
			Str("client_id", body.ClientId).
			Msg("Token generation request received")

		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to decode request body")
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// look up requested user
		var hashPassword string

		row := db.QueryRow("SELECT client_password FROM Clients WHERE client_id = $1", body.ClientId)

		err = row.Scan(&hashPassword)

		log.Debug().
			Str("client_id", body.ClientId).
			Msg("Retrieved hashed password from database")

		if err != nil {
			if err == sql.ErrNoRows {
				log.Error().
					Str("client_id", body.ClientId).
					Msg("Client does not exist in database")
				http.Error(w, "Invalid Client Id", http.StatusBadRequest)
				return
			}
		}
		err = utils.CheckPassword(hashPassword, body.Password)

		if err != nil {
			log.Error().
				Str("client_id", body.ClientId).
				Msg("Invalid password provided")
			http.Error(w, "Invalid Client Id or password", http.StatusUnauthorized)
			return
		}

		log.Info().
			Str("client_id", body.ClientId).
			Msg("Token generated successfully")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Token Generated Successfully",
		})
	}
}
