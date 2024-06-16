package usecase

import (
	"hackathon/dao"
	"log"
)

func Valuation(tweet_id int, sender_user_id int, valuation_type int) error {
	err := dao.Valuation(tweet_id, sender_user_id, valuation_type)
	if err != nil {
		log.Println("failed to Valuation in usecase")
		return err
	}
	return nil
}
