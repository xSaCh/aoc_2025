package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`)

	result := part1(input)
	require.Equal(t, 1227775554, result)
}
func Test2(t *testing.T) {
	input := strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`)

	result := part2(input)
	require.Equal(t, 4174379265, result)
}

func TestRepeatingDigits(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "1234", expected: false},
		{input: "123321", expected: false},
		{input: "123123", expected: true},
		{input: "2121212118", expected: false},
		{input: "2121212119", expected: false},
		{input: "2121212121", expected: true},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, c.expected, hasRepeatingPattern(c.input))
		})
	}
}

func BenchmarkPart1(b *testing.B) {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	by, _ := io.ReadAll(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rd := strings.NewReader(string(by))
		part1(rd)
	}
}
func BenchmarkPart2(b *testing.B) {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	by, _ := io.ReadAll(f)

	for b.Loop() {
		rd := strings.NewReader(string(by))
		part2(rd)
	}
}
