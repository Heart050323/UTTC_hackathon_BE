package model

type PostTweet struct {
	SenderUserID   int    `json:"sender_user_id"`
	Content        string `json:"content"`
	RepliedTweetID int    `json:"replied_tweet_id"`
	ReTweetID      int    `json:"re_tweet_id"`
}

type RetweetOnResponse struct {
	RetweetOn int `json:"re_tweet_on"`
}

type RetweetOnRequest struct {
	TweetID      int `json:"tweet_id"`
	SenderUserID int `json:"sender_user_id"`
}
