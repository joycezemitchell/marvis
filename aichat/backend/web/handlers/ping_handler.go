package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func (h *ChatHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Return a success response
	json.NewEncoder(w).Encode(&Response{
		Message: "ok",
	})
}
