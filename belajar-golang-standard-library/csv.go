package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvStrings := "eko,kurniawan,khannedy\nbudi,pratama,setiawan\njoko,widodo,susanto"

	reader := csv.NewReader(strings.NewReader(csvStrings))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println("error:", err)
			break
		}
		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"eko", "kurniawan", "khannedy"})
	writer.Write([]string{"budi", "pratama", "setiawan"})
	writer.Write([]string{"joko", "widodo", "susanto"})
	writer.Flush()
}