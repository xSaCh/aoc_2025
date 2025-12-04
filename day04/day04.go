package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	fmt.Println(part1(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	grid := [][]bool{}
	for sc.Scan() {
		line := sc.Text()
		row := make([]bool, len(line))
		for i, l := range line {
			if l == '@' {
				row[i] = true
			}
		}
		grid = append(grid, row)
	}
	// fmt.Printf("grid: %v\n", grid)
	// fmt.Printf("getNumAdjustantNeigbours(grid, 2, 0): %v\n", getNumAdjustantNeigbours(grid, 2, 0))
	count := 0
	for r := range grid {
		for c := range grid {
			if grid[r][c] {
				neighbours := getNumAdjustantNeigbours(grid, c, r)
				if neighbours < 4 {
					count++
				}

			}
		}
	}

	return count
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	grid := [][]bool{}
	for sc.Scan() {
		line := sc.Text()
		row := make([]bool, len(line))
		for i, l := range line {
			if l == '@' {
				row[i] = true
			}
		}
		grid = append(grid, row)
	}
	// fmt.Printf("grid: %v\n", grid)
	// fmt.Printf("getNumAdjustantNeigbours(grid, 2, 0): %v\n", getNumAdjustantNeigbours(grid, 2, 0))
	count := 0
	for r := range grid {
		for c := range grid {
			if grid[r][c] {
				neighbours := getNumAdjustantNeigbours(grid, c, r)
				if neighbours < 4 {
					count++
				}

			}
		}
	}

	return count
}

func getNumAdjustantNeigbours(grid [][]bool, x, y int) int {
	num := 0
	n := len(grid[0]) // width n height are same
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if y+r >= n || x+c >= n || (r == c && r == 0) || y+r < 0 || x+c < 0 {
				continue
			}
			// _, isSelected := selected[[2]int{y + r, x + c}]
			if grid[y+r][x+c] {
				num++
			}
		}
	}
	return num
}
