package main

import "fmt"

func main(){
	const firstname = "wilson"
	const lastname string = "william"

	fmt.Println(firstname)

	const (
		age int = 23
		is_done = false
	)

	fmt.Println(is_done)
}