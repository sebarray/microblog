package model

type Tweet struct {
	ID      string `json:"id" bson:"_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	Content string `json:"content" bson:"content"`
}
