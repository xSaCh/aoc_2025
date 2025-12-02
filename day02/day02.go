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
	sc.Scan()
	productIds := strings.Split(sc.Text(), ",")

	invalidIds := make(map[int]struct{}, 0)
	for _, p := range productIds {
		idParts := strings.Split(p, "-")
		firstId, _ := strconv.Atoi(idParts[0])
		lastId, _ := strconv.Atoi(idParts[1])

		firstFHalf, firstLHalf := halfStringToInts(idParts[0])
		lastFHalf, lastLHalf := halfStringToInts(idParts[1])

		if firstFHalf >= firstId && firstFHalf <= lastId {
			invalidIds[firstFHalf] = struct{}{}
		}
		if firstLHalf >= firstId && firstLHalf <= lastId {
			invalidIds[firstLHalf] = struct{}{}
		}
		if lastFHalf >= firstId && lastFHalf <= lastId {
			invalidIds[lastFHalf] = struct{}{}
		}
		if lastLHalf >= firstId && lastLHalf <= lastId {
			invalidIds[lastLHalf] = struct{}{}
		}
	}
	fmt.Printf("invalidIds: %v\n", invalidIds)
	sum := 0
	for id := range invalidIds {
		sum += id
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

func halfStringToInts(strA string) (int, int) {
	fHalf := strA[:len(strA)/2]
	sHalf := strA[len(strA)/2:]
	intfHalf, _ := strconv.Atoi(fHalf + fHalf)
	intsHalf, _ := strconv.Atoi(sHalf + sHalf)

	return intfHalf, intsHalf
}
