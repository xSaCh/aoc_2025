package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	fmt.Println(part1(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	reds := []Coord{}
	for sc.Scan() {
		point := strings.Split(sc.Text(), ",")
		xs, _ := strconv.Atoi(point[0])
		ys, _ := strconv.Atoi(point[1])

		reds = append(reds, Coord{xs, ys})
	}

	slices.SortFunc(reds, func(a, b Coord) int {
		if a.X > b.X {
			return 1
		} else if a.X < b.X {
			return -1
		}
		return 0
	})
	maxArea := 0
	for i := 0; i < len(reds)-2; i++ {
		j := len(reds) - 1
		for j != 0 {
			a := reds[i]
			b := reds[j]
			area := calculateArea(a, b)
			maxArea = max(maxArea, area)
			j--
		}
	}
	return maxArea
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	reds := []Coord{}
	for sc.Scan() {
		point := strings.Split(sc.Text(), ",")
		xs, _ := strconv.Atoi(point[0])
		ys, _ := strconv.Atoi(point[1])

		reds = append(reds, Coord{xs, ys})
	}

	slices.SortFunc(reds, func(a, b Coord) int {
		if a.X > b.X {
			return 1
		} else if a.X < b.X {
			return -1
		}
		return 0
	})
	fmt.Printf("reds: %v\n", reds)
	maxArea := 0
	for i := 0; i < len(reds)-2; i++ {
		j := len(reds) - 1
		for j != 0 {
			a := reds[i]
			b := reds[j]
			area := calculateArea(a, b)
			maxArea = max(maxArea, area)
			j--
		}
	}
	return maxArea
}

func calculateArea(a, b Coord) int {
	l := math.Abs(float64(a.X)-float64(b.X)) + 1
	h := math.Abs(float64(a.Y)-float64(b.Y)) + 1
	ar := l * h
	if ar < 0 {
		return -int(ar)
	}
	return int(ar)
}
