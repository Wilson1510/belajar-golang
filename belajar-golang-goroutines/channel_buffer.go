package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "Hello"
	fmt.Println(<-ch)
	fmt.Println(cap(ch))
	fmt.Println(len(ch))
	ch <- "World"
	ch <- "Golang"
	ch <- "Channel"
	fmt.Println(cap(ch))
	fmt.Println(len(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}