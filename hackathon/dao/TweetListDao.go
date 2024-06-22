package dao

import (
	"fmt"
	"hackathon/model"
	"log"
	"time"
)

func TweetList() ([]model.TweetData, error) {
	rows, err := db.Query(`SELECT 
	tweet.tweet_id, tweet.content, tweet.replied_tweet_id, tweet.re_tweet_id, tweet.created_at, tweet.likecount, tweet.badcount, tweet.replycount, tweet.re_tweetcount,
	user.user_name, user.user_id
	FROM tweet
	JOIN user ON tweet.sender_user_id = user.user_id
	WHERE tweet.replied_tweet_id = 0
	ORDER BY tweet.created_at DESC`)
	if err != nil {
		log.Println("tweetlistのDBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.BadCount, &tweetData.ReplyCount, &tweetData.ReTweetCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return nil, err
		}
		tweetData.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			log.Println("Failed to parse created_at:", err)
			return nil, err
		}

		tweetDatas = append(tweetDatas, tweetData)
	}
	return tweetDatas, nil
}

func ReplyTweetList(replied_tweet_id int) ([]model.TweetData, error) {
	rows, err := db.Query(`SELECT 
	tweet.tweet_id, tweet.content, tweet.replied_tweet_id, tweet.re_tweet_id, tweet.created_at, tweet.likecount, tweet.badcount, tweet.replycount, tweet.re_tweetcount,
	user.user_name, user.user_id
	FROM tweet
	JOIN user ON tweet.sender_user_id = user.user_id
	WHERE tweet.replied_tweet_id = ?
	ORDER BY tweet.created_at DESC`, replied_tweet_id)
	if err != nil {
		log.Println("replied tweet list DBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.BadCount, &tweetData.ReplyCount, &tweetData.ReTweetCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return nil, err
		}
		tweetData.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			log.Println("Failed to parse created_at:", err)
			return nil, err
		}

		tweetDatas = append(tweetDatas, tweetData)
	}

	return tweetDatas, nil
}

func TweetCall(tweet_id int) (model.TweetData, error) {
	rows, err := db.Query(`SELECT 
        tweet.tweet_id, tweet.content, tweet.replied_tweet_id, tweet.re_tweet_id, tweet.created_at, tweet.likecount, tweet.badcount, tweet.replycount, tweet.re_tweetcount,
        user.user_name, user.user_id
        FROM tweet
        JOIN user ON tweet.sender_user_id = user.user_id
        WHERE tweet.tweet_id = ?`, tweet_id)
	if err != nil {
		log.Println("Tweet Call DBクエリが叩けてません:", err)
		return model.TweetData{}, err
	}
	defer rows.Close()

	var tweetData model.TweetData
	if rows.Next() {
		var createdAt string
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &createdAt, &tweetData.LikeCount, &tweetData.BadCount, &tweetData.ReplyCount, &tweetData.ReTweetCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
		if err != nil {
			log.Println("Scan failed:", err)
			return model.TweetData{}, err
		}
		tweetData.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			log.Println("Failed to parse created_at:", err)
			return model.TweetData{}, err
		}
	} else {
		log.Println("No tweet found with tweet_id:", tweet_id)
		return model.TweetData{}, fmt.Errorf("no tweet found with tweet_id: %d", tweet_id)
	}

	if rows.Next() {
		log.Println("Warning: multiple tweets found with the same tweet_id:", tweet_id)
	}

	return tweetData, nil
}

func PastTweetList(user_id int) ([]model.TweetData, error) {
	rows, err := db.Query(`SELECT 
	tweet.tweet_id, tweet.content, tweet.replied_tweet_id, tweet.re_tweet_id, tweet.created_at, tweet.likecount, tweet.badcount, tweet.replycount, tweet.re_tweetcount,
	user.user_name, user.user_id
	FROM tweet
	JOIN user ON tweet.sender_user_id = user.user_id
	WHERE user.user_id = ?
	ORDER BY tweet.created_at DESC`, user_id)
	if err != nil {
		log.Println("PasttweetlistのDBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.BadCount, &tweetData.ReplyCount, &tweetData.ReTweetCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return nil, err
		}
		tweetData.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			log.Println("Failed to parse created_at:", err)
			return nil, err
		}

		tweetDatas = append(tweetDatas, tweetData)
	}
	return tweetDatas, nil
}

func LikeTweetList(user_id int) ([]model.TweetData, error) {
	rows, err := db.Query(`SELECT 
		tweet.tweet_id, 
		tweet.content, 
		tweet.replied_tweet_id, 
		tweet.re_tweet_id, 
		tweet.created_at, 
		tweet.likecount, 
		tweet.badcount, 
		tweet.replycount, 
		tweet.re_tweetcount,
		user.user_name, 
		user.user_id
	FROM tweet
	JOIN valuation ON tweet.tweet_id = valuation.tweet_id
	JOIN user ON tweet.sender_user_id = user.user_id
	WHERE tweet.sender_user_id = ? AND valuation.valuation_type = 1
	ORDER BY tweet.created_at DESC`, user_id)
	if err != nil {
		log.Println("liketweetlistのDBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.BadCount, &tweetData.ReplyCount, &tweetData.ReTweetCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return nil, err
		}
		tweetData.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			log.Println("Failed to parse created_at:", err)
			return nil, err
		}

		tweetDatas = append(tweetDatas, tweetData)
	}
	return tweetDatas, nil
}

func BadTweetList(user_id int) ([]model.TweetData, error) {
	rows, err := db.Query(`SELECT 
		tweet.tweet_id, 
		tweet.content, 
		tweet.replied_tweet_id, 
		tweet.re_tweet_id, 
		tweet.created_at, 
		tweet.likecount, 
		tweet.badcount, 
		tweet.replycount, 
		tweet.re_tweetcount,
		user.user_name, 
		user.user_id
	FROM tweet
	JOIN valuation ON tweet.tweet_id = valuation.tweet_id
	JOIN user ON tweet.sender_user_id = user.user_id
	WHERE tweet.sender_user_id = ? AND valuation.valuation_type = 2
	ORDER BY tweet.created_at DESC`, user_id)
	if err != nil {
		log.Println("liketweetlistのDBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.BadCount, &tweetData.ReplyCount, &tweetData.ReTweetCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return nil, err
		}
		tweetData.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			log.Println("Failed to parse created_at:", err)
			return nil, err
		}

		tweetDatas = append(tweetDatas, tweetData)
	}
	return tweetDatas, nil
}
