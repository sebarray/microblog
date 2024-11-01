package repository

import "microblog/internal/command/model"

type CommandRepositoryInterface interface {
	SaveTweet(tweet *model.Tweet) error
}
