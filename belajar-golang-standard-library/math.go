package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-1000))
	fmt.Println(math.Round(10.5))
	fmt.Println(math.Ceil(10.3))
	fmt.Println(math.Floor(10.7))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))

	fmt.Println(math.Sin(1.57))
}