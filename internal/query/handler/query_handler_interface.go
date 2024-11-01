package handler

import "net/http"

type CommandHandlerInterface interface {
	PostTweet(w http.ResponseWriter, r *http.Request)
	FollowUser(w http.ResponseWriter, r *http.Request)
}
