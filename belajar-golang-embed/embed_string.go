package main

import (
	_ "embed"
	"fmt"
)

//go:embed manage.txt
var version string

func main() {
	fmt.Println(version)
}