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
	fmt.Println(part1(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)

	freshIdRanges := make([][2]int, 0)
	availableIds := []int{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			break
		}
		ids := strings.Split(line, "-")
		a, _ := strconv.Atoi(ids[0])
		b, _ := strconv.Atoi(ids[1])
		freshIdRanges = append(freshIdRanges, [2]int{a, b})

	}
	for sc.Scan() {
		id := sc.Text()
		a, _ := strconv.Atoi(id)
		availableIds = append(availableIds, a)
	}
	freshIds := []int{}
	count := 0
	for _, id := range availableIds {
		for _, idRange := range freshIdRanges {
			if id >= idRange[0] && id <= idRange[1] {
				freshIds = append(freshIds, id)
				count++
				break
			}
		}
	}
	// fmt.Printf("freshIdRanges: %v\n", freshIdRanges)
	// fmt.Printf("availableIds: %v\n", availableIds)
	// fmt.Printf("freshIds: %v\n", freshIds)
	return count
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		_ = sc.Text()

	}
	return 0
}
