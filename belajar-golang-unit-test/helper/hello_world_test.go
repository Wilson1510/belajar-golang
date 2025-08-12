package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Andi")
	if result != "Hello Andi" {
		panic("Result is not Hello Andi")
	}
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		panic("Result is not Hello Budi")
	}
}