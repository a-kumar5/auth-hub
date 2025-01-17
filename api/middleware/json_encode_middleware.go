package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func JsonEncoderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().
			Str("path", r.URL.Path).
			Msg("Setting Content-Type header to application/json")

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)

		log.Debug().
			Str("path", r.URL.Path).
			Msg("Completed JSON encoding middleware")
	})
}
