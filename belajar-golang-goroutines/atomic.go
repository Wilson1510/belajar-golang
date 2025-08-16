package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", counter)
}