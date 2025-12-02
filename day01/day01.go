package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)

	STARTING_POINT := 50
	curPoint := STARTING_POINT
	reachedZero := 0

	for sc.Scan() {
		rotationsLine := sc.Text()
		dir := rotationsLine[0]
		distance, _ := strconv.Atoi(rotationsLine[1:])
		switch dir {
		case 'L':
			curPoint = (curPoint - distance) % 100
		case 'R':
			curPoint = (curPoint + distance) % 100
		}
		if curPoint == 0 {
			reachedZero++
		}
	}
	return reachedZero
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)

	STARTING_POINT := 50
	curPoint := STARTING_POINT
	reachedZero := 0

	for sc.Scan() {
		rotationsLine := sc.Text()
		dir := rotationsLine[0]
		distance, _ := strconv.Atoi(rotationsLine[1:])
		for range distance {
			switch dir {
			case 'L':
				curPoint = curPoint - 1
			case 'R':
				curPoint = curPoint + 1
			}
			curPoint = mod(curPoint)
			if curPoint == 0 {
				reachedZero++
			}
		}
	}
	return reachedZero
}

func mod(a int) int {
	return (a%100 + 100) % 100
}

func quotient(a int) float64 {
	return math.Floor(float64(a) / float64(100))
}
