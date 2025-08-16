package main

import "fmt"

func fungsi(m string) {
	m = "Nilai diubah di dalam fungsi!"
	fmt.Println(m)
}

func main() {
    // Variabel global (atau setidaknya di dalam fungsi main)
    message := "Hello dari main!"

    // Ini adalah IIFE
    func(m string) {
        // 'm' adalah variabel lokal yang hanya ada di dalam fungsi ini
        // Nilainya adalah salinan dari 'message' yang diberikan saat pemanggilan
        m = "Nilai diubah di dalam IIFE!" // Perubahan ini hanya berlaku di sini
        fmt.Println(m)
    }(message) // Memanggil fungsi anonim dan meneruskan 'message'

    // 'message' di sini tidak terpengaruh oleh perubahan di dalam IIFE
    fmt.Println(message)
}