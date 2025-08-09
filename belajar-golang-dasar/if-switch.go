package main

import "fmt"

func main(){
	name := "kurniawan"

	if name == "budi" {
		fmt.Println("Hello budi")
	} else if name == "john" {
		fmt.Println("Hello john")
	} else {
		fmt.Println("Hello semuanya")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	age := 26

	switch age {
	case 23:
		fmt.Println("umur 23")
	case 24:
		fmt.Println("umur 24")
	default:
		fmt.Println("umur tidak ditemukan")
	}
} 
