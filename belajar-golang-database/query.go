package main

import (
	"context"
	"database/sql"
	"os"
	"fmt"
	"time"
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

	script := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		fmt.Println("email: ", email.String)
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		fmt.Println("birth_date: ", birth_date.Time)
		fmt.Println("married: ", married)
		fmt.Println("created_at: ", created_at)
		fmt.Println("--------------------------------")
	}
}