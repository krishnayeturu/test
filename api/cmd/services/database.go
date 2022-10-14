package services

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() {
	USERNAME := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DRIVER := os.Getenv("DRIVER")
	if DRIVER == "" {
		DRIVER = "mysql"
	}
	dbconf := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DB_NAME + "?charset=utf8mb4" + "&parseTime=True"
	db, err := sql.Open(DRIVER, dbconf)
	if err != nil {
		log.Panicf("Error connecting to database : error=%v", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	// return db, err
	Db = db
}

func CloseDB() error {
	return Db.Close()
}
