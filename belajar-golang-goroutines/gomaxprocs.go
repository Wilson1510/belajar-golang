package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			time.Sleep(3 * time.Second)
			wg.Done()
		}()
	}
	
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)
	wg.Wait()

	runtime.GOMAXPROCS(20)
	totalThread = runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)
	
}