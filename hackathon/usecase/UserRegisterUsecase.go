package usecase

import (
	"hackathon/dao"
	"log"
)

func UserRegister(email string, user_name string) error {
	err := dao.UserRegister(email, user_name)
	if err != nil {
		log.Println("failed to user register in usecase")
		return err
	}
	return nil
}
