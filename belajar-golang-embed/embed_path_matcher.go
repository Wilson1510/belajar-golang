package main

import (
	"embed"
	"fmt"
)

//go:embed *
var path embed.FS

func main() {
	a, _ := path.ReadFile("note.txt")
	b, _ := path.ReadFile("manage.txt")
	c, _ := path.ReadFile("embed_string.go")
	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}