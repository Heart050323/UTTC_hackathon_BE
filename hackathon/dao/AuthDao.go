package dao

import (
	"hackathon/model"
	"log"
	"time"
)

func Auth(email string) (*model.UserInfo, error) {
	rows, err := db.Query("SELECT u.user_id, u.user_name, t.tweet_id, t.content, t.replied_tweet_id, t.re_tweet_id, t.created_at FROM user u JOIN tweet t ON u.user_id = t.sender_user_id WHERE u.email = ?", email)
	if err != nil {
		log.Println("DBクエリが叩けてません")
		return nil, err
	}
	defer rows.Close()

	var userInfo model.UserInfo
	tweetDatas := make([]model.TweetData, 0)

	for rows.Next() {
		var tweetData model.TweetData
		var created_at string
		err := rows.Scan(&userInfo.User_id, &userInfo.User_Name, &tweetData.Tweet_id, &tweetData.Content, &tweetData.RepliedTweetID, &tweetData.ReTweetID, &created_at)
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

	userInfo.TweetDatas = tweetDatas
	return &userInfo, nil
}
