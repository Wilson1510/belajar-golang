 
package main

import "fmt"

func main(){
	var makanan [4]int
	makanan[0] = 98
	makanan[2] = 70

	fmt.Println(makanan[0])
	fmt.Println(makanan[1])
	fmt.Println(makanan[2])
	fmt.Println(makanan[3])

	var umur = [...]string {"jus", "kopi", "teh", "kurma"}
	fmt.Println(umur)
	fmt.Println(len(umur))
	fmt.Println(umur[3])
}