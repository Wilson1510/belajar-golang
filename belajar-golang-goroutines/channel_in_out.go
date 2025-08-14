package main

import (
	"fmt"
	"time"
)


func sendMessage(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Hello, World!"
}

func receiveMessage(ch <-chan string) {
	message := <-ch
	fmt.Println(message)
}

func main() {
	ch := make(chan string)

	go sendMessage(ch)
	go receiveMessage(ch)

	fmt.Println("Done")

	time.Sleep(3 * time.Second)
	fmt.Println("Closing channel")
	close(ch)
}