package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func PastTweet(email string) (*model.WholeData, error) {
	wholeData, err := dao.PastTweet(email)
	if err != nil {
		log.Println("failed to wholedata call in usecase")
		return nil, err
	}
	return wholeData, nil
}
