package model

import "time"

type EmailRequest struct {
	Email string `json:"email"`
}

type WholeData struct {
	User_id    int         `json:"user_id"`
	User_Name  string      `json:"user_name"`
	TweetDatas []TweetData `json:"tweet_datas"`
}
type TweetData struct {
	Tweet_id       int       `json:"tweet_id"`
	Content        string    `json:"content"`
	RepliedTweetID int       `json:"replied_tweet_id"`
	ReTweetID      int       `json:"re_tweet_id"`
	CreatedAt      time.Time `json:"created_at"`
}
