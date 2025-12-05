package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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
	count := 0
	for _, id := range availableIds {
		for _, idRange := range freshIdRanges {
			if id >= idRange[0] && id <= idRange[1] {
				count++
				break
			}
		}
	}
	return count
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)

	freshIdRanges := make([][2]int, 0)
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
	// Ignoring available ids, Irrelevant for part 2

	slices.SortFunc(freshIdRanges, func(x, y [2]int) int {
		if x[0] < y[0] {
			return -1
		} else if x[0] > y[0] {
			return 1
		}
		return 0
	})
	mergeIdRanges := make([][2]int, 0)
	for _, idRange := range freshIdRanges {

		if len(mergeIdRanges) != 0 && idRange[0] <= mergeIdRanges[len(mergeIdRanges)-1][1] {
			mergeIdRanges[len(mergeIdRanges)-1][1] =
				max(mergeIdRanges[len(mergeIdRanges)-1][1], idRange[1])
		} else {
			mergeIdRanges = append(mergeIdRanges, idRange)
		}
	}
	count := 0
	for _, idRange := range mergeIdRanges {
		count += idRange[1] - idRange[0] + 1
	}
	return count
}
