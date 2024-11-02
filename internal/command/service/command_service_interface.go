package service

import "microblog/internal/command/model"

type CommandServiceInterface interface {
	CreateTweet(tweet *model.Tweet) error
	FollowUser(followerID, followeeID string) error
}
