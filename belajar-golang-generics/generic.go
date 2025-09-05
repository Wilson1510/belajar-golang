package main

import (
	"fmt"
)

type Bag[T any] []T

func PrintBag[E any](bag Bag[E]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func main() {
	PrintBag[int](Bag[int]{1, 2, 3})
	PrintBag[string](Bag[string]{"Hello", "World"})
	PrintBag[bool](Bag[bool]{true, false})
}

