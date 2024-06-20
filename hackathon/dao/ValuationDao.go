package dao

import (
	"hackathon/model"
	"log"
)

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
	if valuation_type == 1 || valuation_type == -1 {
		_, err = tx.Exec("UPDATE tweet SET likecount = likecount + ? WHERE tweet_id = ?;", valuation_type, tweet_id)
		log.Println("goodcountをいじりました")
		if err != nil {
			tx.Rollback()
			log.Println("failed to set likecount")
			return err
		}
	}
	if valuation_type == 2 || valuation_type == -2 {
		_, err = tx.Exec("UPDATE tweet SET badcount = badcount + ? WHERE tweet_id = ?;", valuation_type/2, tweet_id)
		log.Println("badcountをいじりました")
		if err != nil {
			tx.Rollback()
			log.Println("failed to set badcount")
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println("failed to commit transaction")
		return err
	}
	return nil
}

func GetValuationType(tweet_id int, sender_user_id int) (model.ValuationTypeResponse, error) {
	rows, err := db.Query("SELECT valuation_type FROM valuation WHERE tweet_id = ? AND sender_user_id = ?", tweet_id, sender_user_id)
	if err != nil {
		log.Println("valuationのDBクエリが叩けてません")
		return model.ValuationTypeResponse{}, err
	}
	defer rows.Close()
	var ValuationType model.ValuationTypeResponse
	for rows.Next() {
		err := rows.Scan(&ValuationType.ValuationType)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return model.ValuationTypeResponse{}, err
		}
	}
	return ValuationType, nil
}
