package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Andi")
	require.Equal(t, "Hello Andi", result, "Result is not Hello Andi")
	fmt.Println("TestSkip Done")
}

func TestHelloWorldAndi(t *testing.T) {
	result := HelloWorld("Andi")
	if result != "Hello Andi" {
		t.Error("Result is not Hello Andi")
	}
	assert.Equal(t, "Hello Andi", result, "Result is not Hello Andi")
	fmt.Println("TestHelloWorldAndi Done")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	// if result != "Hello Budi" {
	// 	t.Fatal("Result is not Hello Budi")
	// }
	require.Equal(t, "Hello Budi", result, "Result is not Hello Budi")
	fmt.Println("TestHelloWorldBudi Done")
}