package fastprint

import (
	"bytes"
	"strings"
	"testing"
)

func TestSprintS(t *testing.T) {
	result := SprintS("abc", "def", "ghijklmno")
	if result != "abcdefghijklmno" {
		t.Error("concat result error: ", result)
	}
}

func BenchmarkSprintS(b *testing.B) {
	var s string
	input := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, "xddd")
	}
	b.ResetTimer()
	s = SprintS(input...)
	s = s
}

func BenchmarkStringsJoin(b *testing.B) {
	input := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, "xddd")
	}
	b.ResetTimer()
	c := ""
	strings.Join(input, c)
}

func BenchmarkStringPlusLoop(b *testing.B) {
	input := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, "xddd")
	}
	b.ResetTimer()
	c := ""
	for _, i := range input {
		c += i
	}
}

func BenchmarkBprintS(b *testing.B) {
	input := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, "xddd")
	}
	b.ResetTimer()
	BprintS(input...)
}

func BenchmarkSprintB(b *testing.B) {
	var s string
	input := make([][]byte, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, []byte("xddd"))
	}
	b.ResetTimer()
	s = SprintB(input...)
	s = s
}

func BenchmarkBprintB(b *testing.B) {
	input := make([][]byte, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, []byte("xddd"))
	}
	b.ResetTimer()
	BprintB(input...)
}

func BenchmarkBytesJoin(b *testing.B) {
	input := make([][]byte, 0, b.N)
	for i := 0; i < b.N; i++ {
		input = append(input, []byte("xddd"))
	}
	b.ResetTimer()
	c := []byte{}
	bytes.Join(input, c)
}
