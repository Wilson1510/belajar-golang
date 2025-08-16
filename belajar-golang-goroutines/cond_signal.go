package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	queue := make([]int, 0)

	// Goroutine Produsen
	go func() {
		for i := 1; i <= 5; i++ {
			mu.Lock()
			queue = append(queue, i)
			fmt.Printf("Produsen: Menambahkan item %d ke antrean\n", i)
			mu.Unlock()
			cond.Signal() // Memberi tahu salah satu goroutine konsumen bahwa ada data baru
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Goroutine Konsumen
	for i := 1; i <= 2; i++ {
		go func(id int) {
			for {
				mu.Lock()
				// Cek kondisi: apakah antrean kosong?
				for len(queue) == 0 {
					fmt.Printf("Konsumen %d: Antrean kosong, menunggu...\n", id)
					cond.Wait() // Melepaskan mutex dan menunggu sinyal
				}
				// Saat dibangungkan, mutex sudah terkunci kembali
				item := queue[0]
				queue = queue[1:]
				fmt.Printf("Konsumen %d: Mengambil item %d\n", id, item)
				mu.Unlock()
			}
		}(i)
	}

	// Biarkan program berjalan sebentar
	time.Sleep(5 * time.Second)
	fmt.Println("Program selesai.")
}