package bootstrap

import "log"

type Application struct {
	Env      *Env
	Postgres *Postgres
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	log.Println("Closing DB Connection")
	ClosePostgresDBConnection(*app.Postgres)
}
