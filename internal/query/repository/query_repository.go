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

// GetFollowedTweets returns the tweets of the users followed by the given userID.
func (r *QueryRepository) GetFollowedTweets(userID string) ([]model.Tweet, error) {
	// Paso 1: Obtener IDs de los usuarios seguidos por el userID
	followersCollection := r.db.Collection("followers")
	cursor, err := followersCollection.Find(context.TODO(), bson.M{"follower_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Crear una lista de IDs de los usuarios seguidos
	var followedIDs []string
	for cursor.Next(context.TODO()) {
		var follow model.Follow
		if err := cursor.Decode(&follow); err != nil {
			return nil, err
		}
		followedIDs = append(followedIDs, follow.FollowedID)
	}

	// Paso 2: Obtener los tweets de los usuarios seguidos
	tweetsCollection := r.db.Collection("tweets")
	tweetsCursor, err := tweetsCollection.Find(context.TODO(), bson.M{"user_id": bson.M{"$in": followedIDs}})
	if err != nil {
		return nil, err
	}
	defer tweetsCursor.Close(context.TODO())

	// Decodificar los tweets encontrados
	var tweets []model.Tweet
	for tweetsCursor.Next(context.TODO()) {
		var tweet model.Tweet
		if err := tweetsCursor.Decode(&tweet); err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
