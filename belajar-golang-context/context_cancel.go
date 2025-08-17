package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine berhenti karena context dibatalkan")
			return
		default:
			fmt.Println("Melakukan pekerjaan...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(runtime.NumGoroutine())
	go doWork(ctx)
	fmt.Println(runtime.NumGoroutine())

	// Biarkan goroutine berjalan sebentar
	time.Sleep(2 * time.Second)

	// Batalkan context untuk menghentikan goroutine
	cancel()
	fmt.Println(runtime.NumGoroutine())
	// Beri waktu goroutine untuk selesai
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}