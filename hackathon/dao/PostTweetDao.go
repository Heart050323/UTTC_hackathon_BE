package dao

import (
	"hackathon/model"
	"log"
)

func PostTweet(sender_user_id int, content string, replied_tweet_id int, re_tweet_id int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to begin transaction")
		return err
	}
	_, err = tx.Exec("INSERT INTO tweet (sender_user_id, content, replied_tweet_id ,re_tweet_id) VALUES (?,?,?,?)", sender_user_id, content, replied_tweet_id, re_tweet_id)
	if err != nil {
		tx.Rollback()
		log.Println("Failed to insert tweet")
		return err
	}
	if replied_tweet_id != 0 {
		_, err = tx.Exec("UPDATE tweet SET replycount = replycount + 1 WHERE tweet_id = ?", replied_tweet_id)
		if err != nil {
			tx.Rollback()
			log.Println("Failed to update replycount")
			return err
		}
	}
	if re_tweet_id != 0 {
		_, err = tx.Exec("UPDATE tweet SET re_tweetcount = re_tweetcount + 1 WHERE tweet_id = ?", re_tweet_id)
		if err != nil {
			tx.Rollback()
			log.Println("Failed to update retweetcount")
			return err
		}
		_, err = tx.Exec("INSERT INTO retweeton (sender_user_id, tweet_id, re_tweet_on) VALUES (?,?,1)", sender_user_id, re_tweet_id)
		if err != nil {
			tx.Rollback()
			log.Println("Failed to retweetOn")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction")
		return err
	}
	return nil
}

func GetRetweetOn(tweet_id int, sender_user_id int) (model.RetweetOnResponse, error) {
	rows, err := db.Query(`SELECT re_tweet_on FROM retweeton WHERE tweet_id = ? AND sender_user_id = ? `, tweet_id, sender_user_id)
	if err != nil {
		log.Println("RetweerOn DBクエリが叩けてません:", err)
		return model.RetweetOnResponse{}, err
	}
	defer rows.Close()

	var RetweetOnResponse model.RetweetOnResponse
	if rows.Next() {
		err := rows.Scan(&RetweetOnResponse.RetweetOn)
		if err != nil {
			log.Println("Scan failed:", err)
			return model.RetweetOnResponse{}, err
		}
	}
	return RetweetOnResponse, nil
}
