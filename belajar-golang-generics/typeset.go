package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum[int](1, 2))
	fmt.Println(sum[float64](1.5, 2.2))
	fmt.Println(sum[string]("Hello", "World"))
}


type Number interface {
	int | float64 | string
}

func sum[T Number](a T, b T) T {
	if a < b {
		return b
	} else {
		return a
	}
}

