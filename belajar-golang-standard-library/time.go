package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	date := time.Date(2025, 8, 11, 10, 10, 10, 10, time.Local)
	fmt.Println(date)

	utc := time.Date(2025, 8, 11, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	parse, _ := time.Parse(time.RFC3339, "2025-08-11T10:10:10Z")
	fmt.Println(parse)

	formatter := "2006-01-02 15:04:05"
	parse, _ = time.Parse(formatter, "2025-08-31 10:10:10")
	fmt.Println(parse)

	var duration1 time.Duration = time.Hour * 24
	fmt.Println(duration1)

	duration2 := time.Minute * 10
	fmt.Println(duration2)

	duration3 := time.Second * 10
	fmt.Println(duration3)

	duration4 := time.Millisecond * 100
	fmt.Println(duration4)

	duration5 := time.Microsecond * 100
	fmt.Println(duration5)

	total := duration1 + duration2 + duration3 + duration4 + duration5
	fmt.Println(total)

	fmt.Printf("Total: %d\n", total)
	
}