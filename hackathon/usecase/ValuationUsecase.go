package usecase

import (
	"hackathon/dao"
	"hackathon/model"
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

func GetValuationType(tweet_id int, sender_user_id int) (model.ValuationTypeResponse, error) {
	ValuationType, err := dao.GetValuationType(tweet_id, sender_user_id)
	if err != nil {
		log.Println("failed to ValuationTypeRequest in usecase")
		return model.ValuationTypeResponse{}, err
	}
	return ValuationType, nil
}
