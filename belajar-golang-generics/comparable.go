package main

import (
	"fmt"
)

func isSame[T comparable](value1 T, value2 T) bool {
	return value1 == value2
}

func main() {
	fmt.Println(isSame(1, 1))
	fmt.Println(isSame(1, 2))
	fmt.Println(isSame("John", "John"))
	fmt.Println(isSame("John", "Doe"))
	fmt.Println(isSame(true, true))
	fmt.Println(isSame(true, false))
}