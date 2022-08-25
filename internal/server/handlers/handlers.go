package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// SendResponse оборачивает ответ в JSON формат
func SendResponse(w http.ResponseWriter, res any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal(err)
	}
}

// SendError оборачивает ошибку в JSON формат
func SendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	jErr := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	if jErr != nil {
		log.Fatal(jErr)
	}
}
