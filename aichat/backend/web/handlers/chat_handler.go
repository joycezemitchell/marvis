package handlers

import (
	"aichat/internal/app/auth"
	"aichat/internal/app/chat"
	"aichat/internal/app/dto"
	"encoding/json"
	"net/http"
)

type ChatHandler struct {
	service     chat.Service
	authService auth.AuthService
}

func NewChatHandler(s chat.Service, a auth.AuthService) *ChatHandler {
	return &ChatHandler{service: s, authService: a}
}

func (h *ChatHandler) HandleChat(w http.ResponseWriter, r *http.Request) {
	var req dto.ChatRequest
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.ProcessChat(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
	// TODO: write the customers to w as JSON or HTML
}
