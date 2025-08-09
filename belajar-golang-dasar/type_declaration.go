 package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool

	var contoh NoKTP = "111222"
	var contoh2 = "555333"
	var marriedStatus Married = true
	var contohKtp NoKTP = NoKTP(contoh2)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
	fmt.Println(marriedStatus)

}