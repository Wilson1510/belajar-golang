package main

import (
	"context"
	"fmt"
)

func main() {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")
	ctxF := context.WithValue(ctxC, "f", "F")

	ctxG := context.WithValue(ctxD, "g", "G")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)
	fmt.Println(ctxG)

	fmt.Println("\n")
	fmt.Println(ctxG.Value("b"))
	fmt.Println(ctxG.Value("d"))
	fmt.Println(ctxG.Value("g"))

	fmt.Println("\n")
	fmt.Println(ctxF.Value("c"))
	fmt.Println(ctxF.Value("b"))

	fmt.Println("\n")
	fmt.Println(ctxE.Value("a"))
	fmt.Println(ctxE.Value("e"))

	fmt.Println("\n")
	fmt.Println(ctxA.Value("a"))
	fmt.Println(ctxA.Value("c"))
}