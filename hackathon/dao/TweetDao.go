package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

var db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("No .env file, %v\n", err)
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	// ①-2
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlUserPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}
func PostTweet(sender_user_id int, content string, replied_tweet_id int, re_tweet_id int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to begin transaction")
		return err
	}
	_, err = tx.Exec("INSERT INTO tweet (sender_user_id, content, replied_tweet_id ,re_tweet_id) VALUES (?,?,?,?)", sender_user_id, content, replied_tweet_id, re_tweet_id)
	if err != nil {
		tx.Rollback()
		log.Println("Failed to insert tweet")
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction")
		return err
	}
	return nil
}
