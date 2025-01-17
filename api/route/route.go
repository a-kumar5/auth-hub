package route

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/a-kumar5/auth-hub/api/controller"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func RegisterRoutes(router *mux.Router) {
	log.Info().Msg("Registering base routes")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Handling request to root endpoint")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Welcome to auth-hub")
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Handling health check request")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Service is up and running")
	})

	log.Info().Msg("Base routes registered successfully")
}

func RegisterClientRoutes(router *mux.Router, db *sql.DB) {
	log.Info().Msg("Registering client routes")

	router.HandleFunc("/client", controller.GetClients(db)).Methods("GET")
	router.HandleFunc("/client", controller.CreateClient(db)).Methods("POST")
	router.HandleFunc("/auth/token", controller.CreateToken(db)).Methods("POST")

	log.Info().Msg("Client routes registered successfully")
}
