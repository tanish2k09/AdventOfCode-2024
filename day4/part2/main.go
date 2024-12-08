package main

import (
	"fmt"
	"os"
	"strings"
)

func readFileFromArg() string {
	if len(os.Args) < 2 {
		panic("Usage: go run 1.go <path>")
	}
	filePath := os.Args[1]

	// Get the filename then read from it
	data, err := os.ReadFile(filePath)
	check(err)
	return string(data)
}

func main() {
	// Make sure we have a CLI path arg
	data := readFileFromArg()
	dataStr := strings.TrimSpace(string(data))
	rows := strings.Split(dataStr, "\n")

	// Process the XMAS strings
	xmasCount := processXmas(rows)

	// We're done!
	fmt.Println(xmasCount)
}

func processDiagonal(grid []string, xPos [2]int) int {
	y, x := xPos[0], xPos[1]

	if !((grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S') ||
		(grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M')) {
		return 0
	}

	if !((grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S') ||
		(grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M')) {
		return 0
	}

	return 1
}

func processXmas(grid []string) int {
	total := 0

	for y := 1; y < len(grid)-1; y++ {
		row := grid[y]
		for x := 1; x < len(row)-1; x++ {
			if row[x] == 'A' {
				coords := [2]int{y, x}
				total += processDiagonal(grid, coords)
			}
		}
	}

	return total
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
