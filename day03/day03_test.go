package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`)

	result := part1(input)
	require.Equal(t, 357, result)
}
func Test2(t *testing.T) {
	input := strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`)

	result := part2(input)
	require.Equal(t, 0, result)
}
