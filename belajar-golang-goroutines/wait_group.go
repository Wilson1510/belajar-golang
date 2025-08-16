package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Pastikan Done() dipanggil saat goroutine selesai
	defer wg.Done()

	fmt.Printf("Worker %d: Mulai...\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d: Selesai!\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Menambah penghitung WaitGroup
		go worker(i, &wg)
	}

	// Menunggu sampai semua goroutine selesai (penghitung mencapai 0)
	wg.Wait()
	fmt.Println("Semua worker telah selesai. Program utama berakhir.")
}