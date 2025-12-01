package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)

	result := part1(input)
	require.Equal(t, 3, result)
}
func Test2(t *testing.T) {
	input := strings.NewReader(``)

	result := part2(input)
	require.Equal(t, 0, result)
}
