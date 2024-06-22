package model

import "time"

type EmailRequest struct {
	Email string `json:"email"`
}

type UserIDRequest struct {
	User_id int `json:"user_id"`
}

type ReplyRequest struct {
	RepliedTweetID int `json:"replied_tweet_id"`
}

type TweetRequest struct {
	TweetID int `json:"tweet_id"`
}

type TweetData struct {
	SenderUserID   int       `json:"user_id"`
	SenderUserName string    `json:"user_name"`
	TweetID        int       `json:"tweet_id"`
	Content        string    `json:"content"`
	RepliedTweetID int       `json:"replied_tweet_id"`
	ReTweetID      int       `json:"re_tweet_id"`
	CreatedAt      time.Time `json:"created_at"`
	LikeCount      int       `json:"likecount"`
	BadCount       int       `json:"badcount"`
	ReplyCount     int       `json:"replycount"`
	ReTweetCount   int       `json:"re_tweetcount"`
}
