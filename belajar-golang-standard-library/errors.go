package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivideByZero = errors.New("Pembagi tidak boleh 0")
	ErrInvalidNilai = errors.New("Nilai tidak boleh 0")
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, ErrDivideByZero
	}
	if nilai == 0 {
		return 0, ErrInvalidNilai
	}
	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagian(7, 1)
	if err != nil {
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("Pembagi tidak boleh 0:", err.Error())
		} else if errors.Is(err, ErrInvalidNilai) {
			fmt.Println("Nilai tidak boleh 0:", err.Error())
		} else{
			fmt.Println("Error:", err.Error())
		}
	} else {
		fmt.Println("Hasil:", hasil)
	}
}