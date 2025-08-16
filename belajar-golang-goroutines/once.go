package main

import (
	"fmt"
	"sync"
)	

var counter = 0

func onlyOnce() {
	counter++
	fmt.Println("Only Once")
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			once.Do(onlyOnce)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}