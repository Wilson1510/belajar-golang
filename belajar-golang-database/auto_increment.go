package main

import (
	"database/sql"
	"fmt"
	"os"
	"context"
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
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, script, "test@test.com", "test")
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Insert Id: ", lastInsertId)
}