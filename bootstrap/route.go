package bootstrap

import (
	"encoding/json"
	"net/http"

	"github.com/a-kumar5/auth-hub/api/controller"
	"github.com/a-kumar5/auth-hub/api/middleware"

	"github.com/rs/zerolog/log"
)

func (app *Application) registerRoutes() {
	log.Info().Msg("Registering base routes")

	app.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Handling request to root endpoint")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Welcome to auth-hub")
	})

	app.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Handling health check request")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Service is up and running")
	})

	log.Info().Msg("Base routes registered successfully")
}

func (app *Application) registerApplicationRoutes() {
	log.Info().Msg("Registering client routes")
	r := app.Router
	r.Path("/api/v1/auth/token").Handler(http.HandlerFunc(controller.CreateToken(app.Postgres.SQLDB)))
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.ValidateAuth)
	api.Path("/client").Handler(controller.GetClients(app.Postgres.SQLDB)).Methods("GET")
	api.Path("/client").Handler(controller.CreateClient(app.Postgres.SQLDB)).Methods("POST")
	log.Info().Msg("Client routes registered successfully")
}
