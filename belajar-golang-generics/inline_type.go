package main

import "fmt"

func FindMin[T interface{int | float64 | int64}](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func GetFirst[T []E, E any](data T) E {
	return data[0]
}

func main() {
	fmt.Println(FindMin(1, 2))
	fmt.Println(FindMin(1.5, 2.2))
	fmt.Println(FindMin(1, 2.2))

	fmt.Println(GetFirst([]int{1, 2, 3, 4, 5}))
	fmt.Println(GetFirst([]float64{1.5, 2.2, 3.3, 4.4, 5.5}))
	fmt.Println(GetFirst([]string{"Hello", "World", "Go", "Generics"}))
}