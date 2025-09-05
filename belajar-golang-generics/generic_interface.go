package main

import "fmt"

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[E any](getterSetter GetterSetter[E], value E) E {
	getterSetter.SetValue(value)
	return getterSetter.GetValue()
}

type MyData[S any] struct {
	Value S
}

func (d *MyData[W]) GetValue() W {
	return d.Value
}

func (d *MyData[D]) SetValue(value D) {
	d.Value = value
}

func main() {
	fmt.Println(ChangeValue[int](&MyData[int]{Value: 1}, 2))
	fmt.Println(ChangeValue[string](&MyData[string]{Value: "Hello"}, "World"))
	fmt.Println(ChangeValue[bool](&MyData[bool]{Value: true}, false))
}