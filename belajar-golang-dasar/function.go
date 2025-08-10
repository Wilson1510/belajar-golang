 package main

import "fmt"


type Filter func(string) string

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}


func personAge(name string, age int8){
	fmt.Println("Hello", name, "umur", age)
}

func personStatus(status string) string {
	return status
}

func getFullName() (string, string){
	return "wilson", "john"
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}


func biodata()(name string, age string, status string){
	name = "wilson"
	age = "23"
	status = "single"
	return age, status, name
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}


func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func logging(){
	fmt.Println("Selesai memanggil function")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApplication(error bool){
	defer logging()
	fmt.Println("Run Application")
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main(){
	personAge("wilson", 23)
	personAge("budi", 24)
	personAge("john", 25)

	var a string = personStatus("single")
	fmt.Println(a)
	b := personStatus("married")
	fmt.Println(b)

	firstName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(getFullName())

	a, b, c := biodata()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)
	numbers := []int {10, 20, 30, 40, 50}
	total = sumAll(numbers...)
	fmt.Println(total)
	total = sumAll()
	fmt.Println(total)
	total = sumAll(20, 45, 10)
	fmt.Println(total)

	contoh:= getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("wilson"))
	fmt.Println(misal("budi"))

	sayHelloWithFilter("wilson", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
	sayHelloWithFilter("budi", spamFilter)

	sayHelloWithFilter2("budi", spamFilter)
	sayHelloWithFilter2("anjing", spamFilter)
	sayHelloWithFilter2("wilson", spamFilter)

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("wilson", blacklist)
	registerUser("anjing", blacklist)

	registerUser("budi", func(name string) bool {
		return name == "anjing"
	})

	fmt.Println(factorialRecursive(10))

	counter := 0
	increment := func() {
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)

	runApplication(false)

	// ini adalah contoh komentar
	/**sini juga bisa
	ini adlaah komentar banyak baris
	bisa tambah terus
	tambah
	tambah
	*/

}