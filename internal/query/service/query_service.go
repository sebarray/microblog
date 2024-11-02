package service

import (
	"microblog/internal/query/model"
	"microblog/internal/query/repository"
)

type QueryService struct {
	repo repository.QueryRepositoryInterface
}

// NewQueryService creates a new CommandService.
func NewQueryService(repo repository.QueryRepositoryInterface) *QueryService {
	return &QueryService{repo: repo}
}

// GetFollowedTweets implements QueryServiceInterface.
func (s QueryService) GetFollowedTweets(userID string) ([]model.Tweet, error) {

	//todo: verificar que el usuario exista
	return s.repo.GetFollowedTweets(userID)
}
