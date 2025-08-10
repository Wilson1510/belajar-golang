package main

import (
	"belajar-golang-dasar/helper"
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	helper.SayHello("John")
	fmt.Println(database.Connect())
}

