package repository

import "microblog/internal/query/model"

type QueryRepositoryInterface interface {
	GetFollowedTweets(userID string) ([]model.Tweet, error)
}
