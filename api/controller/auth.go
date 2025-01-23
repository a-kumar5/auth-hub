package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/a-kumar5/auth-hub/api/utils"
	_ "github.com/a-kumar5/auth-hub/docs"
	"github.com/rs/zerolog/log"
)

// CreateToken godoc
// @Summary Generate authentication token
// @Description Creates a JWT token for a client using their client_id and password
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body utils.CreateTokenRequest true "Authentication credentials"
// @Success 200 {object} utils.CreateTokenResponse "Token generated successfully"
// @Failure 400 {object} utils.ErrorRes "Invalid request payload or client ID"
// @Failure 401 {object} utils.ErrorRes "Invalid client password"
// @Failure 502 {object} utils.ErrorRes "Token generation failed"
// @Router /auth/token [post]
func CreateToken(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")

		var request utils.CreateTokenRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		log.Info().
			Str("client_id", request.ClientId).
			Msg("Token generation request received")

		if err != nil {
			response := utils.ErrorRes{Message: "Invalid request payload"}
			utils.WriteJSONError(w, response, http.StatusBadRequest)
			log.Error().
				Err(err).
				Msg("Failed to decode request body")
			return
		}

		// look up requested user
		var hashPassword string

		row := db.QueryRow("SELECT client_password FROM Clients WHERE client_id = $1", request.ClientId)

		err = row.Scan(&hashPassword)

		log.Debug().
			Str("client_id", request.ClientId).
			Msg("Retrieved hashed password from database")

		if err != nil {
			if err == sql.ErrNoRows {
				response := utils.ErrorRes{Message: "Invalid Client Id"}
				utils.WriteJSONError(w, response, http.StatusBadRequest)
				log.Error().
					Err(err).
					Msg("Client does not exist in database")
				return
			}
		}
		err = utils.CheckPassword(hashPassword, request.Password)

		if err != nil {
			response := utils.ErrorRes{Message: "Invalid client password"}
			utils.WriteJSONError(w, response, http.StatusUnauthorized)
			log.Error().
				Err(err).
				Msg("Invalid password provided")
			return
		}
		token, err := utils.CreateToken(request.ClientId, secretKey)
		if err != nil {
			response := utils.ErrorRes{Message: "couldn't generate jwt token"}
			utils.WriteJSONError(w, response, http.StatusBadGateway)
			log.Error().
				Err(err).
				Msg("couldn't generate key ")
			return
		}

		log.Info().
			Str("client_id", request.ClientId).
			Msg("Token generated successfully")

		w.Header().Set("Authorization", token)
		response := utils.CreateTokenResponse{
			Message: "Token Generated Successfully",
			Token:   token,
		}
		json.NewEncoder(w).Encode(response)
	}
}
