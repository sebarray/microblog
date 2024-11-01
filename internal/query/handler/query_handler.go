package handler

import (
	"microblog/internal/query/repository"
	"net/http"
)

type QueryHandler struct {
	repo repository.QueryRepository
}

func NewQueryHandler(repo repository.QueryRepository) *QueryHandler {
	return &QueryHandler{repo: repo}
}

func (h *QueryHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	// Lógica para obtener la línea de tiempo
}
