package main

import "fmt"

func main(){
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("\n")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	fmt.Println("\n")

	names := []string {"wilson", "budi", "john"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	fmt.Println("\n")

	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("\n")

	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
	fmt.Println("\n")

	for i := 0; i < 10; i++ {
		if i % 2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
