package handler

import (
	"encoding/json"
	"microblog/internal/command/model"
	"microblog/internal/command/service"
	"net/http"
)

type CommandHandler struct {
	Service service.CommandServiceInterface
}

func NewCommandHandler(Service service.CommandServiceInterface) *CommandHandler {
	return &CommandHandler{Service: Service}
}

// PostTweet creates a new tweet.
func (h *CommandHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	var tweet model.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateTweet(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// FollowUser follows a user.
func (h *CommandHandler) FollowUser(w http.ResponseWriter, r *http.Request) {

	var follow model.Follow
	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.FollowUser(follow.FollowerID, follow.FollowedID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
