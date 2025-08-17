package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
    if err != nil {
        panic(err)
    }
	defer db.Close()

	fmt.Println("Successfully connected to the database")
}