package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time1 := <-timer.C
	fmt.Println(time1)

	timer_after := time.After(5 * time.Second)
	fmt.Println(timer_after)

	time_after := <-timer_after
	fmt.Println(time_after)

	time.AfterFunc(5 * time.Second, func() {
		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())
	time.Sleep(6 * time.Second)
}