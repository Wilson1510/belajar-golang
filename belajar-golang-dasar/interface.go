package main

import (
	"fmt"
)

// Definisi interface
type bentuk interface {
	luas() float64
}

// Tipe data lingkaran
type lingkaran struct {
	jariJari float64
}

// Tipe data persegi
type persegi struct {
	sisi float64
}

// Implementasi method luas() untuk lingkaran
func (l lingkaran) luas() float64 {
	return 3.14 * l.jariJari * l.jariJari
}

// Implementasi method luas() untuk persegi
func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

// Fungsi yang menerima interface bentuk
func hitungLuas(b bentuk) {
	fmt.Println("Luasnya adalah:", b.luas())
}

func main() {
	ling := lingkaran{jariJari: 5}
	pers := persegi{sisi: 10}

	hitungLuas(ling) // Memanggil dengan tipe lingkaran
	hitungLuas(pers) // Memanggil dengan tipe persegi

	var i interface{}
    i = 10         // i sekarang menampung integer
    fmt.Println(i)

    i = "Halo, Go!"  // i sekarang menampung string
    fmt.Println(i)

    i = struct{
        nama string
    }{
        nama: "Gemini",
    } // i sekarang menampung struct anonim
    fmt.Println(i)
}