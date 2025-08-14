package main

import (
    "fmt"
    "time"
)

func ambilData1(ch chan string) {
    time.Sleep(2 * time.Second) // simulasi proses lambat
    ch <- "Data dari sumber 1"
}

func ambilData2(ch chan string) {
    time.Sleep(1 * time.Second)
    ch <- "Data dari sumber 2"
}

func main() {
    ch := make(chan string)

    go ambilData1(ch)
    go ambilData2(ch)

    // Menerima data dari dua goroutine
    for i := 0; i < 2; i++ {
        fmt.Println(<-ch)
    }

    fmt.Println("Selesai")
}
