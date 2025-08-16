package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	start := time.Now()
	counter := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", counter)
	fmt.Println("Time:", time.Since(start))

}