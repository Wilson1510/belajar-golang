package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}


func main(){
	var budi Customer
	fmt.Println(budi)
	budi.Name = "budi"
	budi.Address = "jakarta"
	budi.Age = 23
	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Address)
	fmt.Println(budi.Age)

	joko := Customer{
		Name: "joko",
		Address: "bandung",
		Age: 24,
	}
	fmt.Println(joko)

	john := Customer{"john", "surabaya", 25}
	fmt.Println(john)

	budi.sayHello("joko")
	

} 
