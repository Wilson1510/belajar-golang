package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAndi(t *testing.T) {
	result := HelloWorld("Andi")
	if result != "Hello Andi" {
		t.Error("Result is not Hello Andi")
	}
	fmt.Println("TestHelloWorldAndi Done")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		t.Fatal("Result is not Hello Budi")
	}
	fmt.Println("TestHelloWorldBudi Done")
}