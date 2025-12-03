package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	fmt.Println(part2(f))
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

func part2(input io.Reader) uint64 {
	sc := bufio.NewScanner(input)
	var sum uint64 = 0

	for sc.Scan() {
		bank := sc.Text()
		from := 0
		joltage := strings.Builder{}
		for tailSpace := range 12 {
			right := 12 - 1 - tailSpace
			a := bank[from : len(bank)-right]
			maxLeft, maxLeftIdx := myMax(a)

			from += maxLeftIdx + 1
			joltage.WriteRune(maxLeft)
		}
		j, _ := strconv.Atoi(joltage.String())
		sum += uint64(j)
	}
	return sum
}

func myMax(bank string) (rune, int) {
	maxLeft := '0'
	maxLeftIndex := 0
	for i, b := range bank {
		if b > maxLeft {
			maxLeft = b
			maxLeftIndex = i
		}
	}
	return maxLeft, maxLeftIndex
}
