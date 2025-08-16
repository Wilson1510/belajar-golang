package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    workers := 3 // Jumlah pekerja

    for i := 1; i <= workers; i++ {
        go func(id int) {
            mu.Lock()
            defer mu.Unlock()
            fmt.Printf("Pekerja %d: Menunggu sinyal untuk mulai.\n", id)
            
            // Masuk ke mode menunggu
            cond.Wait()
            
            // Setelah dibangungkan oleh Broadcast(), mutex terkunci kembali
            fmt.Printf("Pekerja %d: Menerima sinyal, mulai bekerja!\n", id)
        }(i)
    }

    // Biarkan pekerja memiliki waktu untuk masuk ke dalam mode menunggu
    time.Sleep(1 * time.Second)

    fmt.Println("\nMain: Mengirim sinyal 'broadcast' untuk membangunkan semua pekerja.")
    
    // Broadcast: Membangunkan semua goroutine yang menunggu
    mu.Lock()
    cond.Broadcast()
    mu.Unlock()

    // Biarkan pekerja memiliki waktu untuk menyelesaikan pekerjaan mereka
    time.Sleep(1 * time.Second)
    fmt.Println("\nMain: Program selesai.")
}