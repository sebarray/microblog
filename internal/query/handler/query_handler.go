package handler

import (
	"encoding/json"
	"microblog/internal/query/service"
	"net/http"
)

type QueryHandler struct {
	Service service.QueryServiceInterface
}

func NewQueryHandler(Service service.QueryServiceInterface) *QueryHandler {
	return &QueryHandler{Service: Service}
}

func (h *QueryHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {

	//QUERY PARAM
	userID := r.URL.Query().Get("userID")

	tweets, err := h.Service.GetFollowedTweets(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonTweets, err := json.Marshal(tweets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonTweets)

}
