package main

import (
	"fmt"
	"maps"

	"golang.org/x/exp/constraints"
)

func FindMin[T constraints.Ordered](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func main() {
	fmt.Println(FindMin(1, 2))
	fmt.Println(FindMin(1.5, 2.2))
	fmt.Println(FindMin(1, 2.2))
	fmt.Println(FindMin("Hello", "World"))

	map1 := map[string]string{
		"name": "John",
		"age": "20",
	}
	map2 := map[string]string{
		"name": "Jane",
		"age": "21",
	}

	fmt.Println(maps.Equal(map1, map2))
}
