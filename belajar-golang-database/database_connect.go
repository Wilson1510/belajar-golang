package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {
	dbUser := os.Getenv("password")
	db, err := sql.Open("mysql", "root:"+dbUser+"@tcp(localhost:3306)/golang")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

func main() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES ('buku', 'buku')"

	result, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	fmt.Println("Successfully connected to the database")
}