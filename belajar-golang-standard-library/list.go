package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	var tail *list.Element = data.Back()
	fmt.Println(tail.Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}