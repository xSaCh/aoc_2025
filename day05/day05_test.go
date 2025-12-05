package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	result := part1(input)
	require.Equal(t, 3, result)
}
func Test2(t *testing.T) {
	input := strings.NewReader(``)

	result := part2(input)
	require.Equal(t, 0, result)
}
