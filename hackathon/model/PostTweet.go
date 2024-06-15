package model

type PostTweet struct {
	SenderUserID   int    `json:"sender_user_id"`
	Content        string `json:"content"`
	RepliedTweetID int    `json:"replied_tweet_id"`
	ReTweetID      int    `json:"re_tweet_id"`
}
