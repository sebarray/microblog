package repository

import "microblog/internal/query/model"

type QueryRepositoryInterface interface {
	GetTweetsByUser(userID string) ([]model.Tweet, error)
}
