package bootstrap

/*
const (
  host     = "localhost"
  port     = 5433
  user     = "postgres"
  password = "your-password"
  dbname   = "calhounio_demo"
)
*/

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
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

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		postgresURI = fmt.Sprintf("postgres://%s:%s?sslmode=disable", dbHost, dbPort)
	}

	sqlDB, err := sql.Open("postgres", postgresURI)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("database/sql connection established")

	return &Postgres{
		SQLDB: sqlDB,
	}
}

func ClosePostgresDBConnection(db Postgres) {

	err := db.SQLDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to Postgres DB closed.")
}
