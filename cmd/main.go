package main

import (
	"log"

	"github.com/a-kumar5/auth-hub/bootstrap"
)

func main() {
	log.Println("Hello, Welcome to auth-hub service")
	app := bootstrap.App()

	env := app.Env
	db := app.Postgres.SQLDB
	defer app.CloseDBConnection()

	log.Println(env, db)
}
