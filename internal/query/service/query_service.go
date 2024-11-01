package service

import (
	"microblog/internal/query/model"
	"microblog/internal/query/repository"
)

type CommandService struct {
	repo *repository.QueryRepository
}

// NewCommandService creates a new CommandService.
func NewCommandService(repo *repository.QueryRepository) *CommandService {
	return &CommandService{repo: repo}
}

// GetFollowedTweets implements QueryServiceInterface.
func (s CommandService) GetFollowedTweets(userID string) ([]model.Tweet, error) {

	//todo: verificar que el usuario exista
	return s.repo.GetFollowedTweets(userID)
}
