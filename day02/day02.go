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
	sc.Scan()
	productIds := strings.Split(sc.Text(), ",")

	sum := 0

	for _, p := range productIds {
		idParts := strings.Split(p, "-")
		firstId, _ := strconv.Atoi(idParts[0])
		lastId, _ := strconv.Atoi(idParts[1])

		for id := firstId; id <= lastId; id++ {
			strId := strconv.Itoa(id)
			if len(strId)%2 != 0 {
				continue
			}

			mid := len(strId) / 2
			if strId[:mid] == strId[mid:] {
				sum += id
			}
		}
	}

	return sum
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	sc.Scan()
	productIds := strings.Split(sc.Text(), ",")

	sum := 0

	for _, p := range productIds {
		idParts := strings.Split(p, "-")
		firstId, _ := strconv.Atoi(idParts[0])
		lastId, _ := strconv.Atoi(idParts[1])

		for id := firstId; id <= lastId; id++ {
			strId := strconv.Itoa(id)
			if hasRepeatingPattern(strId) {
				sum += id
			}
		}
	}
	return sum
}

func hasRepeatingPattern(str string) bool {
outer:
	for n := 1; n <= len(str)/2; n++ {
		if len(str)%n != 0 {
			continue
		}
		a := str[:n]
		for i := n; i <= len(str)-n; i += n {
			if a != str[i:i+n] {
				continue outer
			}
		}
		return true
	}
	return false
}
