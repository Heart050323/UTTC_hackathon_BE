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
