package dao

import "log"

func Valuation(tweet_id int, sender_user_id int, valuation_type int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to begin transaction")
		return err
	}
	_, err = tx.Exec(`INSERT INTO valuation (tweet_id, sender_user_id, valuation_type) VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE valuation_type = ?,
	created_at = CURRENT_TIMESTAMP`, tweet_id, sender_user_id, valuation_type, valuation_type)
	if err != nil {
		tx.Rollback()
		log.Println("failed to insert Valuation")
		return err
	}
	_, err = tx.Exec("UPDATE tweet SET likecount = likecount + ? WHERE tweet_id = ?;", valuation_type, tweet_id)
	if err != nil {
		tx.Rollback()
		log.Println("failed to set likecount")
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("failed to commit transaction")
		return err
	}
	return nil
}
