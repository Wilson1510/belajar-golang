package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "guest", "Put your username")
	password := flag.String("password", "guest", "Put your password")
	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}