package model

type Follow struct {
	FollowerID string `json:"follower_id" bson:"follower_id"`
	FollowedID string `json:"followed_id" bson:"followed_id"`
}
