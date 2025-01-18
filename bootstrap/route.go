package bootstrap

import (
	"encoding/json"
	"net/http"

	"github.com/a-kumar5/auth-hub/api/controller"
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

func (app *Application) registerClientRoutes() {
	log.Info().Msg("Registering client routes")

	app.Router.HandleFunc("/client", controller.GetClients(app.Postgres.SQLDB, app.Env.SecretKey)).Methods("GET")
	app.Router.HandleFunc("/client", controller.CreateClient(app.Postgres.SQLDB)).Methods("POST")
	app.Router.HandleFunc("/auth/token", controller.CreateToken(app.Postgres.SQLDB, app.Env.SecretKey)).Methods("POST")

	log.Info().Msg("Client routes registered successfully")
}
