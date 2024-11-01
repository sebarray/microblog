package handler

import (
	"encoding/json"
	"microblog/internal/command/model"
	"microblog/internal/command/repository"
	"net/http"
)

type CommandHandler struct {
	repo repository.CommandRepository
}

func NewCommandHandler(repo repository.CommandRepository) *CommandHandler {
	return &CommandHandler{repo: repo}
}

func (h *CommandHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	var tweet model.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.SaveTweet(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CommandHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	// Lógica para seguir a un usuario
}
