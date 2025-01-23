package utils

import (
	"encoding/json"
	"net/http"
)

type CreateTokenResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func WriteJSONError(w http.ResponseWriter, response ErrorRes, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
