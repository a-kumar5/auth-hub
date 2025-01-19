package middleware

import (
	"net/http"
	"os"

	"github.com/a-kumar5/auth-hub/api/utils"
	"github.com/rs/zerolog/log"
)

func ValidateAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")
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
		next.ServeHTTP(w, r)
	})
}
