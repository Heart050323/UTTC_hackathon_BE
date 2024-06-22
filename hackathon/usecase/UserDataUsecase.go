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

func GetUserProfile(user_id int) (model.UserProfile, error) {
	userProfile, err := dao.GetUserProfile(user_id)
	if err != nil {
		log.Println("failed to userProfile call in usecase")
		return model.UserProfile{}, err
	}
	return userProfile, nil
}

func UserProfileModify(user_id int, user_name string, status_message *string) error {
	err := dao.UserProfileModify(user_id, user_name, status_message)
	if err != nil {
		log.Println("failed to userProfileModify call in usecase")
		return err
	}
	return nil
}
