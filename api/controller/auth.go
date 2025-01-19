package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/a-kumar5/auth-hub/api/utils"
	"github.com/rs/zerolog/log"
)

func CreateToken(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")

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
					Err(err).
					Msg("Client does not exist in database")
				http.Error(w, "Invalid Client Id", http.StatusBadRequest)
				return
			}
		}
		err = utils.CheckPassword(hashPassword, body.Password)

		if err != nil {
			log.Error().
				Err(err).
				Msg("Invalid password provided")
			http.Error(w, "Invalid Client Id or password", http.StatusUnauthorized)
			return
		}
		token, err := utils.CreateToken(body.ClientId, secretKey)
		if err != nil {
			log.Error().
				Err(err).
				Msg("couldn't generate key ")
			http.Error(w, "couldn't generate key", http.StatusBadGateway)
			return
		}

		log.Info().
			Str("client_id", body.ClientId).
			Msg("Token generated successfully")

		w.Header().Set("Authorization", token)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Token Generated Successfully",
			"token":   token,
		})
	}
}
