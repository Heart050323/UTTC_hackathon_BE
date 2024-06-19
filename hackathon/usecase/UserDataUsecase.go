package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func GetUserInfo(email string) (model.UserData, error) {
	userData, err := dao.GetUserInfo(email)
	if err != nil {
		log.Println("failed to userdata call in usecase")
		return model.UserData{}, err
	}
	return userData, nil
}
