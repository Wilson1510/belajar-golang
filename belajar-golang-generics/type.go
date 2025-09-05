package main

import (
	"fmt"
)

func Hello[T1 any, T2 any](name T1, name2 T2) {
	fmt.Println("Hello", name)
	fmt.Println("Hello", name2)
}

func main() {
	Hello[string, int]("John", 1)
	Hello[int, string](1, "Doe")
	Hello[bool, int](true, 2)
}
