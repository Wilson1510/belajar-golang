package main

import (
	"database/sql"
	"fmt"
	"os"
	"context"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {
	dbUser := os.Getenv("password")
	db, err := sql.Open("mysql", "root:"+dbUser+"@tcp(localhost:3306)/golang?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	fmt.Println("Start Transaction")
	db := getConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "test" + strconv.Itoa(i) + "@test.com"
		comment := "test" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Last Insert Id: ", lastInsertId)
	}

	err = tx.Commit()
	// err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}