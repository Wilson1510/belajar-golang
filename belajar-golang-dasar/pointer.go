package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func changeName(name *string) {
	*name = "Budi"
}

func main() {
	var name string = "John"
	var namePointer *string = &name

	fmt.Println("Name:", name)
	fmt.Println("Name Pointer:", *namePointer)

	*namePointer = "Jane"
	fmt.Println("Name:", name)
	fmt.Println("Name Pointer:", *namePointer)

	fmt.Println("Address of name:", namePointer)

	angka := new(int)
	fmt.Println("Angka:", *angka)
	fmt.Println("Address of angka:", angka)

	*angka = 10
	fmt.Println("Angka:", *angka)
	fmt.Println("Address of angka:", angka)

	a := "Nama A"
	b := "Nama B"

	fmt.Println("A:", a)
	fmt.Println("B:", b)

	changeName(&a)
	fmt.Println("A:", a)

	man := Man{Name: "John"}
	man.Married()
	fmt.Println(man.Name)
}
