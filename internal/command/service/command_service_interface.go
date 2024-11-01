package service


type CommandRepositoryInterface interface {
	SaveTweet(tweet *model.Tweet) error
	FollowUser(followerID, followeeID string) error
}