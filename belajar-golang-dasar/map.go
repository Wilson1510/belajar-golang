package main

import "fmt"

func main(){
	orang := map[string]string{}
	orang["nama"] = "wilson"
	orang["umur"] = "23"
	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["umur"])

	var orang2 map[string]string = map[string]string{
		"nama": "john",
		"umur": "28",
	}
	fmt.Println(orang2)
	fmt.Println(orang2["nama"])
	fmt.Println(orang2["umur"])
	fmt.Println(len(orang2))
	fmt.Println(orang2["key"])

	delete(orang2, "nama")
	fmt.Println(orang2)
	fmt.Println(len(orang2))

}
