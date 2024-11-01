package service

import (
	"errors"
	"microblog/internal/command/model"
	"microblog/internal/command/repository"
)

type CommandService struct {
	repo repository.CommandRepository
}

func NewCommandService(repo repository.CommandRepository) *CommandService {
	return &CommandService{repo: repo}
}

// CreateTweet implements CommandServiceInterface.
func (s CommandService) CreateTweet(tweet *model.Tweet) error {
	if len(tweet.Content) > 280 {
		return errors.New("tweet content is too long")
	}
	return s.repo.CreateTweet(tweet)
}

// FollowUser implements CommandServiceInterface.
func (s CommandService) FollowUser(followerID string, followeeID string) error {
	return s.repo.FollowUser(followerID, followeeID)

}
