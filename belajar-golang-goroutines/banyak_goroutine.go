package main

import (
	"fmt"
	"time"
)

func printNumbers(number int) {
	fmt.Println("Display ", number)
}

func main() {
	for i := 0; i < 1000000; i++ {
		go printNumbers(i)
	}
}