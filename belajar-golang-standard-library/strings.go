package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("John Doe", "John"))
	fmt.Println(strings.Split("John Doe", " "))
	fmt.Println(strings.ToLower("John Doe"))
	fmt.Println(strings.ToUpper("John Doe"))
	fmt.Println(strings.TrimSpace("  John   Doe  "))
	fmt.Println(strings.Trim("xyz apalah gitu yx", "xy"))
	fmt.Println(strings.ReplaceAll("John Doe John Doe", "John", "Jane"))
}