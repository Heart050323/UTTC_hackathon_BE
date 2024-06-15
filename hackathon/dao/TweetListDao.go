package dao

import (
	"hackathon/model"
	"log"
	"time"
)

func TweetList() ([]model.TweetData, error) {
	rows, err := db.Query(`SELECT 
	tweet.tweet_id, tweet.content, tweet.replied_tweet_id, tweet.re_tweet_id, tweet.created_at, tweet.likecount, 
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
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
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
	tweet.tweet_id, tweet.content, tweet.replied_tweet_id, tweet.re_tweet_id, tweet.created_at, tweet.likecount, 
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
		err := rows.Scan(&tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount, &tweetData.SenderUserName, &tweetData.SenderUserID)
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
func PastTweet(email string) (*model.WholeData, error) {
	rows, err := db.Query("SELECT u.user_id, u.user_name, t.tweet_id, t.content, t.replied_tweet_id, t.re_tweet_id, t.created_at, t.likecount FROM user u JOIN tweet t ON u.user_id = t.sender_user_id WHERE u.email = ?", email)
	if err != nil {
		log.Println("DBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	var wholedata model.WholeData
	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&wholedata.User_id, &wholedata.User_Name, &tweetData.TweetID, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at, &tweetData.LikeCount)
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

	wholedata.TweetDatas = tweetDatas
	return &wholedata, nil
}
