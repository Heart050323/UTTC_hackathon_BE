package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := DB.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

var DB *sql.DB

func init() {
	// ①-1
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
	DB = _db
}
func RegisterUser(id string, userName string, age int) error {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal("Failed to begin transaction")
		return err
	}
	_, err = tx.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", id, userName, age)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to insert user")
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction")
		return err
	}
	return nil
}
