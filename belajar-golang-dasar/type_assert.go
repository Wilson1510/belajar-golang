package main

import "fmt"

func main() {
	var i interface{} = "wahyu"

	switch value := i.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Integer:", value)
	case bool:
		fmt.Println("Boolean:", value)
	default:
		fmt.Println("Unknown type")
	}

	result, ok := i.(int)
	fmt.Println(result)
	fmt.Println(ok)
}
