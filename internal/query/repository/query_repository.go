package repository

import (
	"context"
	"microblog/internal/query/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QueryRepository struct {
	db *mongo.Database
}

func NewQueryRepository(db *mongo.Database) *QueryRepository {
	return &QueryRepository{db: db}
}

func (r *QueryRepository) GetTweetsByUser(userID string) ([]model.Tweet, error) {
	var tweets []model.Tweet
	cursor, err := r.db.Collection("tweets").Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &tweets); err != nil {
		return nil, err
	}
	return tweets, nil
}
