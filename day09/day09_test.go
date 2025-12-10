package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	result := part1(input)
	require.Equal(t, 50, result)
}
func Test2(t *testing.T) {
	input := strings.NewReader(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	result := part2(input)
	require.Equal(t, 24, result)
}
