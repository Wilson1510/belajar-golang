package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/name/test.txt"))
	fmt.Println(filepath.Base("hello/name/test.txt"))
	fmt.Println(filepath.Ext("hello/name/test.txt"))
	fmt.Println(filepath.Join("hello", "name", "test.txt"))
	fmt.Println(filepath.Abs("hello/name/test.txt"))
	fmt.Println(filepath.IsAbs("hello/name/test.txt"))
	fmt.Println(filepath.IsLocal("hello/name/test.txt"))
}