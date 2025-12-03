package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	fmt.Println(part1(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	sum := 0
	for sc.Scan() {
		bank := sc.Text()
		curMax := 0
		for gap := 1; gap < len(bank); gap++ {
			i := 0
			j := gap
			for j < len(bank) {
				combined, _ := strconv.Atoi(string(bank[i]) + string(bank[j]))
				curMax = max(curMax, combined)
				i++
				j++
			}
		}
		sum += curMax
	}
	return sum
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		_ = sc.Text()

	}
	return 0
}
