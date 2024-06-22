package dao

import "log"

func UserRegister(email string, user_name string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to begin transaction")
		return err
	}
	_, err = tx.Exec("INSERT INTO users (email, user_name) VALUES(?,?)", email, user_name)
	if err != nil {
		tx.Rollback()
		log.Println("failed to insert NewUser")
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("failed to commit transaction")
		return err
	}
	return nil
}
