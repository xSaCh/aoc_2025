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

		for id := firstId; id <= lastId; id++ {
			strId := strconv.Itoa(id)
			if len(strId)%2 != 0 {
				continue
			}
			fHalf := strId[:len(strId)/2]
			sHalf := strId[len(strId)/2:]
			if fHalf == sHalf {
				invalidIds[id] = struct{}{}
			}
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
