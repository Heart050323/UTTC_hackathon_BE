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

func TweetCall(tweet_id int) (model.TweetData, error) {
	TweetData, err := dao.TweetCall(tweet_id)
	if err != nil {
		log.Println("failed to tweetdata call in usecase")
		return model.TweetData{}, err
	}
	return TweetData, nil
}

func PastTweetList(user_id int) ([]model.TweetData, error) {
	TweetData, err := dao.PastTweetList(user_id)
	if err != nil {
		log.Println("failed to tweetdata call in usecase")
		return nil, err
	}
	return TweetData, nil
}

func LikeTweetList(user_id int) ([]model.TweetData, error) {
	TweetData, err := dao.LikeTweetList(user_id)
	if err != nil {
		log.Println("failed to tweetdata call in usecase")
		return nil, err
	}
	return TweetData, nil
}

func BadTweetList(user_id int) ([]model.TweetData, error) {
	TweetData, err := dao.BadTweetList(user_id)
	if err != nil {
		log.Println("failed to tweetdata call in usecase")
		return nil, err
	}
	return TweetData, nil
}
