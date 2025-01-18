package bootstrap

import (
	"net/http"

	"github.com/a-kumar5/auth-hub/api/middleware"
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

	return app
}

func (app *Application) InitializeRoutes() {
	app.Router.Use(middleware.AccessLogMiddleware)
	app.Router.Use(middleware.JsonEncoderMiddleware)
	app.registerRoutes()
	app.registerClientRoutes()
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
