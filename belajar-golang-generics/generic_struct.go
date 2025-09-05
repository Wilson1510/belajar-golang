package main

import (
	"fmt"
)

type Data[T any] struct {
	First T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[E]) ChangeSecond(newSecond E) E {
	d.Second = newSecond
	return d.Second
}

func main() {
	data := Data[int]{1, 2}
	fmt.Println(data)

	fmt.Println(data.SayHello("John"))
	fmt.Println(data.ChangeSecond(3))
	fmt.Println(data)
}