package main

import (
	"fmt"
	"strconv"
)

func main() {
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(number)
	}

	numberStr, err := strconv.Atoi("99999")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(numberStr)
	}

	strNumber := strconv.Itoa(1000000)
	fmt.Println(strNumber)

}