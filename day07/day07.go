package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	fmt.Println(part1(f))
}

type Pos struct {
	X int
	Y int
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	grid := [][]bool{}
	startPos := Pos{}
	for sc.Scan() {
		line := sc.Text()
		row := make([]bool, len(line))
		for i, l := range line {
			if l == '^' {
				row[i] = true
			}
			if l == 'S' {
				startPos = Pos{i, len(grid)}
			}
		}
		grid = append(grid, row)
	}
	// fmt.Printf("startPos: %v\n", startPos)

	beams := []Pos{startPos}
	cnt := 0
	for len(beams) != 0 {
		b := beams[0]
		beams = beams[1:]
		b.Y++

		// fmt.Printf("b: %v\n", b)
		if b.Y == len(grid)-1 || b.X < 0 || b.X >= len(grid) {
			continue
		}

		if grid[b.Y][b.X] {
			if !slices.Contains(beams, Pos{b.X - 1, b.Y}) {
				beams = append(beams, Pos{b.X - 1, b.Y})
			}
			if !slices.Contains(beams, Pos{b.X + 1, b.Y}) {
				beams = append(beams, Pos{b.X + 1, b.Y})
			}
			cnt += 1
		} else {
			if !slices.Contains(beams, b) {
				beams = append(beams, b)
			}
		}
	}
	return cnt
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		_ = sc.Text()

	}
	return 0
}
