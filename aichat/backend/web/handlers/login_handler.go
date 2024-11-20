package handlers

import (
	"aichat/internal/app/dto"
	"encoding/json"
	"net/http"
)

func (h *ChatHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.authService.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	json.NewEncoder(w).Encode(resp)
}
