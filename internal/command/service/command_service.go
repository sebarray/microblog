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
	//todo: validar que los IDs no corresponde a un usuario existente
	//todo: validar que el usuaurio no se siga a si mismo
	//todo: validar que el usuario no siga a alguien que ya sigue
	return s.repo.FollowUser(followerID, followeeID)

}
