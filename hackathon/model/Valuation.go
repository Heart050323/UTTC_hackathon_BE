package model

type Valuation struct {
	TweetID       uint `json:"tweet_id"`
	SenderUserID  uint `json:"sender_user_id"`
	ValuationType int  `json:"valuation_type"`
}
