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
			if repeatingDigits(strId) {
				sum += id
			}
		}
	}
	return sum
}

func repeatingDigits(str string) bool {
	for n := 1; n <= len(str)/2; n++ {
		if len(str)%n != 0 {
			continue
		}
		a := str[:n]
		same := true
		for i := n; i <= len(str)-n; i += n {
			b := str[i : i+n]
			if a != b {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}

func halfStringToInts(strA string) (int, int) {
	fHalf := strA[:len(strA)/2]
	sHalf := strA[len(strA)/2:]
	intfHalf, _ := strconv.Atoi(fHalf + fHalf)
	intsHalf, _ := strconv.Atoi(sHalf + sHalf)

	return intfHalf, intsHalf
}
