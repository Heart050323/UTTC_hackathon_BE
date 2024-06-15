package dao

import (
	"hackathon/model"
	"log"
	"time"
)

func PastTweet(email string) (*model.WholeData, error) {
	rows, err := db.Query("SELECT u.user_id, u.user_name, t.tweet_id, t.content, t.replied_tweet_id, t.re_tweet_id, t.created_at FROM user u JOIN tweet t ON u.user_id = t.sender_user_id WHERE u.email = ?", email)
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
		err := rows.Scan(&wholedata.User_id, &wholedata.User_Name, &tweetData.Tweet_id, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at)
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
