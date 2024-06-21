package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func PostTweet(sender_user_id int, content string, replied_tweet_id int, re_tweet_id int) error {
	err := dao.PostTweet(sender_user_id, content, replied_tweet_id, re_tweet_id)
	if err != nil {
		log.Println("failed to PostTweet in usecase")
		return err
	}
	return nil
}

func GetRetweetOn(tweet_id int, sender_user_id int) (model.RetweetOnResponse, error) {
	RetweerOn, err := dao.GetRetweetOn(tweet_id, sender_user_id)
	if err != nil {
		log.Println("failed to retweeton call in usecase")
		return model.RetweetOnResponse{}, err
	}
	return RetweerOn, nil
}
