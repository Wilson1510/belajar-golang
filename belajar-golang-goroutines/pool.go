package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return "Hello"
		},
	}

	pool.Put("World")
	pool.Put("World 2")
	pool.Put("World 3")

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	pool.Put("World 4")


	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}