 
package main

import "fmt"

func main(){
	a := [6]string {"a", "b", "c", "d", "e", "f"}
	fmt.Println(a)

	c := a[1:4]
	fmt.Println(c)

	var d [] string = a[2:]
	fmt.Println(d)

	d[1] = "z"
	fmt.Println(a)

	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:] // 
	fmt.Println(daySlice1)
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "Sabtu lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	fmt.Println("\n")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "wilson"
	newSlice[1] = "budi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "john")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[1] = "budi baru"
	fmt.Println(newSlice)

	fmt.Println("\n")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	copySlice[1] = "Budi copy"
	fmt.Println(copySlice)
	fmt.Println(newSlice)

	iniArray := [...]int {1, 2, 3, 4, 5}
	iniSlice := []int {1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	fmt.Println(cap(iniSlice))

	iniSlice = append(iniSlice, 6)
	fmt.Println(iniSlice)
	fmt.Println(cap(iniSlice))

	iniSlice = append(iniSlice, 7)
	fmt.Println(cap(iniSlice))
	fmt.Println(iniSlice)

	
}