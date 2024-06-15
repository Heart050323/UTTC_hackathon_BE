package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func TweetList() ([]model.TweetData, error) {
	tweetList, err := dao.TweetList()
	if err != nil {
		log.Println("failed to tweetlist call in usecase")
		return nil, err
	}
	return tweetList, nil
}

func ReplyTweetList(replied_tweet_id int) ([]model.TweetData, error) {
	repliedTweetList, err := dao.ReplyTweetList(replied_tweet_id)
	if err != nil {
		log.Println("failed to repliedtweetlist call in usecase")
		return nil, err
	}
	return repliedTweetList, nil
}

func PastTweet(email string) (*model.WholeData, error) {
	wholeData, err := dao.PastTweet(email)
	if err != nil {
		log.Println("failed to wholedata call in usecase")
		return nil, err
	}
	return wholeData, nil
}
