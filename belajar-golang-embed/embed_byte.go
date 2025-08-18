package main

import (
	_ "embed"
	"fmt"
)

//go:embed blok.jpg
var blok []byte

func main() {
	fmt.Println(string(blok))
}