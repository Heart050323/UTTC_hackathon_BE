package dao

import (
	"hackathon/model"
	"log"
)

func GetUserInfo(email string) (model.UserData, error) {
	rows, err := db.Query("SELECT user_id, user_name FROM user WHERE email = ?", email)
	if err != nil {
		log.Println("DBクエリが叩けてません")
		return model.UserData{}, err
	}
	defer rows.Close()

	var userData model.UserData

	for rows.Next() {
		err := rows.Scan(&userData.User_id, &userData.User_Name)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return model.UserData{}, err
		}
	}
	return userData, nil
}

func GetUserProfile(user_id int) (model.UserProfile, error) {
	rows, err := db.Query("SELECT user_id, user_name, status_message FROM user WHERE user_id = ?", user_id)
	if err != nil {
		log.Println("DBクエリが叩けてません")
		return model.UserProfile{}, err
	}
	defer rows.Close()

	var userProfile model.UserProfile

	for rows.Next() {
		err := rows.Scan(&userProfile.User_id, &userProfile.User_Name, &userProfile.StatusMessage)
		if err != nil {
			log.Println(rows, err)
			log.Fatal("Scan failed")
			return model.UserProfile{}, err
		}
	}
	return userProfile, nil
}

func UserProfileModify(user_id int, user_name string, status_message *string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to begin transaction")
		return err
	}
	_, err = tx.Exec("UPDATE user SET user_name = ?, status_message = ?  WHERE user_id = ?", user_name, status_message, user_id)
	if err != nil {
		tx.Rollback()
		log.Println("Failed to update userProfile")
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction")
		return err
	}
	return nil
}
