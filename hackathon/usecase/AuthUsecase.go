package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func Auth(email string) (*model.UserInfo, error) {
	userInfo, err := dao.Auth(email)
	if err != nil {
		log.Println("failed to Auth in usecase")
		return nil, err
	}
	return userInfo, nil
}
