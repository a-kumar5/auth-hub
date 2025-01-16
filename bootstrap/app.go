package bootstrap

import (
	"net/http"

	"github.com/a-kumar5/auth-hub/api/middleware"
	"github.com/a-kumar5/auth-hub/api/route"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Router   *mux.Router
	Env      *Env
	Postgres *Postgres
}

func App() *Application {
	app := &Application{}

	log.Info().Msg("Initializing application components")

	app.Router = mux.NewRouter()
	log.Debug().Msg("Router initialized")

	app.Env = NewEnv()
	log.Debug().
		Interface("config", app.Env).
		Msg("Environment configuration loaded")

	app.Postgres = NewPostgresDatabase(app.Env)

	app.initializeRoutes()
	return app
}

func (app *Application) initializeRoutes() {
	app.Router.Use(middleware.AccessLogMiddleware)
	app.Router.Use(middleware.JsonEncoderMiddleware)
	route.RegisterRoutes(app.Router)
	route.RegisterClientRoutes(app.Router, app.Postgres.SQLDB)
}

func (app *Application) Run(addr string) {
	log.Info().Msgf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, app.Router); err != nil {
		log.Error().Msgf("Failed to start server: %v", err)
	}
}

func (app *Application) CloseDBConnection() {
	log.Info().Msg("Closing DB Connection")
	ClosePostgresDBConnection(*app.Postgres)
}
