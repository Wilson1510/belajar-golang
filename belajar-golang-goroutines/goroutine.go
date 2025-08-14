package main

import (
    "fmt"
    "time"
)

func printNumbers(name string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", name, i)
        time.Sleep(500 * time.Millisecond) // simulasi kerja
    }
}

func main() {
    // Menjalankan 3 goroutine
    go printNumbers("Goroutine-1")
    go printNumbers("Goroutine-2")
    go printNumbers("Goroutine-3")

    // Memberi waktu semua goroutine untuk selesai
    time.Sleep(4 * time.Second)
    fmt.Println("Selesai semua goroutine")
}
