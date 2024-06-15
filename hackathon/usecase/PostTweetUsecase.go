package usecase

import (
	"hackathon/dao"
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
