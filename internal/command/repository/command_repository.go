package repository

import (
	"context"
	"microblog/internal/command/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type CommandRepository struct {
	db *mongo.Database
}

func NewCommandRepository(db *mongo.Database) *CommandRepository {
	return &CommandRepository{db: db}
}

func (r *CommandRepository) SaveTweet(tweet *model.Tweet) error {
	_, err := r.db.Collection("tweets").InsertOne(context.TODO(), tweet)
	return err
}
