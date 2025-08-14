package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "satu detik"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "dua detik"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Diterima dari ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Diterima dari ch2:", msg2)
		default:
			fmt.Println("Waiting")
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}