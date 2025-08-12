package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	names := []string{"eko", "budi", "joko"}

	fmt.Println(slices.Contains(numbers, 3))
	fmt.Println(slices.Contains(names, "eko"))
	fmt.Println(slices.Index(names, "budi"))
	fmt.Println(slices.Index(names, "toni"))
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Max(numbers))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))


}