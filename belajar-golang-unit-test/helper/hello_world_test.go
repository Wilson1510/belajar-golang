package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Andi)",
			request: "Andi",
			expected: "Hello Andi",
		},
		{
			name: "HelloWorld(Budi)",
			request: "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Andi", func(t *testing.T) {
		result := HelloWorld("Andi")
		require.Equal(t, "Hello Andi", result, "Result is not Hello Andi")
	})
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result is not Hello Budi")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

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

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Andi")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Andi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Andi")
		}
	})
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "HelloWorld(Andi)",
			request: "Andi",
		},
		{
			name: "HelloWorld(Budi)",
			request: "Budi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}