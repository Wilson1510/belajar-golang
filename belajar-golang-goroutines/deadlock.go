package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1 sync.Mutex
	var mu2 sync.Mutex

	fmt.Println("Memulai program deadlock...")

	// Goroutine 1
	go func() {
		fmt.Println("Goroutine 1: Mencoba mengunci mu1...")
		mu1.Lock()
		defer mu1.Unlock()

		time.Sleep(100 * time.Millisecond) // Memberi waktu Goroutine 2 untuk mengunci mu2

		fmt.Println("Goroutine 1: Mencoba mengunci mu2...")
		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("Goroutine 1: Berhasil mengunci mu1 dan mu2")
	}()

	// Goroutine 2
	go func() {
		fmt.Println("Goroutine 2: Mencoba mengunci mu2...")
		mu2.Lock()
		defer mu2.Unlock()

		time.Sleep(100 * time.Millisecond) // Memberi waktu Goroutine 1 untuk mengunci mu1

		fmt.Println("Goroutine 2: Mencoba mengunci mu1...")
		mu1.Lock()
		defer mu1.Unlock()

		fmt.Println("Goroutine 2: Berhasil mengunci mu2 dan mu1")
	}()

	// Tunggu sebentar untuk melihat deadlock terjadi
	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}