package main

import (
	"embed"
	"fmt"
)

//go:embed note.txt
//go:embed manage.txt
var files embed.FS

func main() {
	a, _ := files.ReadFile("note.txt")
	b, _ := files.ReadFile("manage.txt")
	fmt.Println(string(a))
	fmt.Println(string(b))
}