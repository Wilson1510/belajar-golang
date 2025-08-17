package main

import (
	"context"
	"fmt"
	"time"
)

func fetchSomething(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operasi selesai dengan sukses!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("Operasi dibatalkan karena: %s\n", err)
	}
}

func main() {
	// Buat context dengan timeout 2 detik
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Penting untuk memanggil cancel untuk membersihkan sumber daya

	fmt.Println("Memulai operasi...")
	fetchSomething(ctx)
}