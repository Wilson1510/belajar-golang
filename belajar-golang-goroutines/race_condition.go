package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", counter)
	fmt.Println("Time:", time.Since(start))

}