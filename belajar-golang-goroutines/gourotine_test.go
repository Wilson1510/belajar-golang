package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func runHelloWorld() {
	fmt.Println("Hello from goroutine")
}

func TestCreateGoroutine(t *testing.T) {
	go runHelloWorld()
	fmt.Println("Hello from main")
	time.Sleep(1 * time.Second)
}