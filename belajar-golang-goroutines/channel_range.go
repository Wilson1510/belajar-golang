package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Send data")
			ch <- fmt.Sprintf("Data %d", i)
			fmt.Println("Finish send data", i)
		}
		close(ch)
	}()

	for data := range ch {
		fmt.Println("\nReceive data")
		fmt.Println(data)
		fmt.Println("Finish receive data")
	}

	fmt.Println("Done")
}