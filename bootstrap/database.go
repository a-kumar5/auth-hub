package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Postgres struct {
	SQLDB *sql.DB // Standard database/sql interface
}

func NewPostgresDatabase(env *Env) *Postgres {
	log.Info().Msg("Initializing Postgres database connection")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	log.Debug().
		Str("host", dbHost).
		Str("port", dbPort).
		Str("user", dbUser).
		Msg("Database connection parameters loaded")

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s/authdb?sslmode=disable", dbUser, dbPass, dbHost, dbPort)

	sqlDB, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to open database connection")
		return nil
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to ping database")
		return nil
	}

	log.Info().Msg("Successfully established database connection")

	return &Postgres{
		SQLDB: sqlDB,
	}
}

func ClosePostgresDBConnection(db Postgres) {
	log.Info().Msg("Attempting to close database connection")

	err := db.SQLDB.Close()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to close database connection")
		return
	}

	log.Info().Msg("Successfully closed database connection")
}
