package bootstrap

import (
	"log"
	"net/http"

	"github.com/a-kumar5/auth-hub/api/route"
	"github.com/gorilla/mux"
)

type Application struct {
	Router   *mux.Router
	Env      *Env
	Postgres *Postgres
}

func App() *Application {
	app := &Application{}
	app.Router = mux.NewRouter()
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)

	app.initializeRoutes()
	return app
}

func (app *Application) initializeRoutes() {
	route.RegisterRoutes(app.Router)
}

func (app *Application) Run(addr string) {
	if err := http.ListenAndServe(addr, app.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (app *Application) CloseDBConnection() {
	log.Println("Closing DB Connection")
	ClosePostgresDBConnection(*app.Postgres)
}
