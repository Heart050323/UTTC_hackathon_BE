package model

type Valuation struct {
	TweetID       int `json:"tweet_id"`
	SenderUserID  int `json:"sender_user_id"`
	ValuationType int `json:"valuation_type"`
}
