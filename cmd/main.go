package main

import (
	"github.com/a-kumar5/auth-hub/bootstrap"

	_ "github.com/a-kumar5/auth-hub/docs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Auth Hub API
// @version 1.0
// @description Authentication and Authorization Service API
// @termsOfService http://auth-hub.io/terms/

// @contact.name Auth Hub Support
// @contact.url http://auth-hub.io/support
// @contact.email support@auth-hub.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().
		Timestamp().
		Caller().
		Logger()

	log.Info().Msg("Starting auth-hub service")

	app := bootstrap.App()
	log.Info().Msg("Application bootstrap completed")
	//env := app.Env
	//db := app.Postgres.SQLDB
	defer func() {
		log.Info().Msg("Closing database connection")
		app.CloseDBConnection()
	}()

	app.InitializeRoutes()
	app.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Info().Str("address", "0.0.0.0:8080").Msg("Starting server")
	app.Run("0.0.0.0:8080")
}
