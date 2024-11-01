package service

import (
    "context"
    "microblog/internal/command/repository"
    "microblog/internal/command/models"
)

type CommandService struct {
    repo repository.CommandRepository
}

func NewCommandService(repo repository.CommandRepository) *CommandService {
    return &CommandService{repo: repo}
}

func (s *CommandService) CreateTweet(ctx context.Context, tweet models.Tweet) error {   
	return s.repo.SaveTweet(ctx, tweet)
}


func (s *CommandService) FollowUser(ctx context.Context, followerID, followeeID string) error {
	return s.repo.FollowUser(ctx, followerID, followeeID)
}
