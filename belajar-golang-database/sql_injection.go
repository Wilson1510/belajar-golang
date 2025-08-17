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
	
	username := "admin'; #"
	password := "adminaaa"

	script_berhasil := "SELECT username, password FROM user where username = '"+username+"' and password = '"+password+"' limit 1"
	rows_berhasil, err_berhasil := db.QueryContext(ctx, script_berhasil)
	if err_berhasil != nil {
		panic(err_berhasil)
	}
	defer rows_berhasil.Close()

	if rows_berhasil.Next() {
		var username, password string
		err_berhasil = rows_berhasil.Scan(&username, &password)
		if err_berhasil != nil {
			panic(err_berhasil)
		}
		fmt.Println("Login berhasil")
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
	} else {
		fmt.Println("Login gagal")
	}

	script_gagal := "SELECT username, password FROM user where username = ? and password = ?"
	rows_gagal, err_gagal := db.QueryContext(ctx, script_gagal, username, password)
	if err_gagal != nil {
		panic(err_gagal)
	}
	defer rows_gagal.Close()

	if rows_gagal.Next() {
		var username, password string
		err_gagal = rows_gagal.Scan(&username, &password)
		if err_gagal != nil {
			panic(err_gagal)
		}
		fmt.Println("Login berhasil")
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
	} else {
		fmt.Println("Login gagal")
	}
}