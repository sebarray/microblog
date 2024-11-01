package repository

import "microblog/internal/command/model"

type CommandRepositoryInterface interface {
	CreateTweet(tweet *model.Tweet) error
	FollowUser(followerID, followeeID string) error
}
