package simple

import "github.com/google/wire"

type Foo struct {}
type Bar struct {}

func NewFoo() *Foo {
	return &Foo{}
}

func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}

func InitializeFooBar2() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}