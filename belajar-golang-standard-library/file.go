package main

import (
	"fmt"
	"os"
	"bufio"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	fmt.Println(file)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	fmt.Println(file)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lines string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines += string(line) + "\n"
	}
	return lines, nil
}

func main() {
	createNewFile("test.txt", "iandisnfseo")
	addToFile("test.txt", "\nini adalah baris baru")
	read, err := readFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(read)
}