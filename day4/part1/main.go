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

func processVertical(grid []string, xPos [2]int) int {
	y, x := xPos[0], xPos[1]
	count := 0

	if y-3 >= 0 && grid[y-3][x] == 'S' && grid[y-2][x] == 'A' && grid[y-1][x] == 'M' {
		fmt.Println("top at ", xPos)
		count++
	}

	if y+3 < len(grid) && grid[y+1][x] == 'M' && grid[y+2][x] == 'A' && grid[y+3][x] == 'S' {
		fmt.Println("bottom at ", xPos)
		count++
	}

	return count
}

func processHorizontal(grid []string, xPos [2]int) int {
	y, x := xPos[0], xPos[1]
	count := 0

	if x-3 >= 0 && grid[y][x-3:x] == `SAM` {
		fmt.Println("left at ", xPos)
		count++
	}

	if x+3 < len(grid[y]) && grid[y][x+1:x+4] == `MAS` {
		fmt.Println("right at ", xPos)
		count++
	}

	return count
}

func processDiagonal(grid []string, xPos [2]int) int {
	y, x := xPos[0], xPos[1]
	count := 0
	if y-3 >= 0 && x-3 >= 0 && grid[y-3][x-3] == 'S' && grid[y-2][x-2] == 'A' && grid[y-1][x-1] == 'M' {
		fmt.Println("top left at ", xPos)
		count++
	}
	if y-3 >= 0 && x+3 < len(grid[y]) && grid[y-3][x+3] == 'S' && grid[y-2][x+2] == 'A' && grid[y-1][x+1] == 'M' {
		fmt.Println("top right at ", xPos)
		count++
	}
	if y+3 < len(grid) && x-3 >= 0 && grid[y+3][x-3] == 'S' && grid[y+2][x-2] == 'A' && grid[y+1][x-1] == 'M' {
		fmt.Println("bottom left at ", xPos)
		count++
	}
	if y+3 < len(grid) && x+3 < len(grid[y]) && grid[y+3][x+3] == 'S' && grid[y+2][x+2] == 'A' && grid[y+1][x+1] == 'M' {
		fmt.Println("bottom right at ", xPos)
		count++
	}
	return count
}

func processXmas(grid []string) int {
	total := 0

	for y, row := range grid {
		for x := 0; x < len(row); x++ {
			if row[x] == 'X' {
				coords := [2]int{y, x}
				fmt.Println("Horizontal at ", coords)
				total += processHorizontal(grid, coords)
				fmt.Println("Vertical at ", coords)
				total += processVertical(grid, coords)
				fmt.Println("Diagonal at ", coords)
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
