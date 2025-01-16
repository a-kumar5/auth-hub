package main

import (
	"github.com/rs/zerolog/log"

	"github.com/a-kumar5/auth-hub/bootstrap"
)

func main() {
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

	log.Info().Str("address", "0.0.0.0:8080").Msg("Starting server")
	app.Run("0.0.0.0:8080")
}
