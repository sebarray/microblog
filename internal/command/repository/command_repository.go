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

func (r *CommandRepository) CreateTweet(tweet *model.Tweet) error {
	_, err := r.db.Collection("tweets").InsertOne(context.TODO(), tweet)
	return err
}

func (r *CommandRepository) FollowUser(followerID, followeeID string) error {
	_, err := r.db.Collection("followers").InsertOne(context.TODO(), map[string]string{"follower_id": followerID, "followee_id": followeeID})
	return err
}
