package service

import "microblog/internal/query/model"

type QueryServiceInterface interface {
	GetFollowedTweets(userID string) ([]model.Tweet, error)
}
