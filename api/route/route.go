package route

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/a-kumar5/auth-hub/api/models"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Welcome to auth-hub")
	})
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Service is up and running")
	})
}

func RegisterClientRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/client", models.CreateClient(db)).Methods("POST")
}
