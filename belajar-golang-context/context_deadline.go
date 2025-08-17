package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func doWork(ctx context.Context) {
	fmt.Println("Memulai pekerjaan...")

	// Simulasi pekerjaan yang membutuhkan waktu lama
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Pekerjaan masih berlangsung...")
		case <-ctx.Done():
			// Pekerjaan dihentikan karena deadline sudah lewat
			err := ctx.Err()
			fmt.Printf("Pekerjaan dibatalkan karena: %s\n", err)
			return
		}
	}
}

func main() {
	// Tentukan deadline: 3 detik dari sekarang
	deadline := time.Now().Add(3 * time.Second)
	fmt.Println(runtime.NumGoroutine())

	// Buat context dengan deadline tersebut
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	fmt.Println(runtime.NumGoroutine())
	defer cancel()

	fmt.Printf("Pekerjaan harus selesai sebelum: %s\n", deadline.Format(time.Kitchen))

	// Jalankan pekerjaan
	go doWork(ctx)
	fmt.Println(runtime.NumGoroutine())

	// Tunggu program selesai
	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}