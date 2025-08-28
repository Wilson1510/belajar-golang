package app

import (
	"database/sql"
	"os"
	"time"

	"belajar-golang-restapi/helper"
)

func NewDB() *sql.DB {
	password := os.Getenv("password")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/golang")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}