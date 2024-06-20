package dao

import (
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
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction")
		return err
	}
	return nil
}
