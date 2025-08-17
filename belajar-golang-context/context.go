package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)

	todo := context.TODO()
	fmt.Println(todo)
}