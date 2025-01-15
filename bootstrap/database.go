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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s/authdb?sslmode=disable", dbUser, dbPass, dbHost, dbPort)
	/*
		if dbUser == "" || dbPass == "" {
			postgresURI = fmt.Sprintf("postgres://%s:%s?sslmode=disable", dbHost, dbPort)
		}
	*/

	sqlDB, err := sql.Open("postgres", postgresURI)

	if err != nil {
		log.Error().Msgf("Can't open database connection: ", err)
		return nil
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Error().Msgf("Can't ping to database: ", err)
		return nil
	}

	log.Info().Msg("database/sql connection established")

	return &Postgres{
		SQLDB: sqlDB,
	}
}

func ClosePostgresDBConnection(db Postgres) {

	err := db.SQLDB.Close()
	if err != nil {
		log.Error().Msgf("Not able to close db connection: ", err)
	}

	log.Info().Msg("Connection to Postgres DB closed.")
}
