package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("This is a string\nwith new line\nhmm yeahh")
	reader := bufio.NewReader(input)
	for {
		line, isPrefix, err := reader.ReadLine()
		fmt.Println(isPrefix, err)
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello, World!\n")
	writer.WriteString("Hello, World 2!\n")
	writer.WriteString("Hello, World 3!\n")
	writer.Flush()
}