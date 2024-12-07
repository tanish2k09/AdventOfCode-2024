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

	// Create the initial map state from the rows of data
	rows := strings.Split(dataStr, "\n")
	mapState := createMap(rows)

	// Get the pathset of the guard
	pathset := getGuardPathset(mapState)

	// We're done!
	fmt.Println(len(pathset))
}

type MapState struct {
	width     int
	height    int
	grid      [][]rune
	guard     [2]int
	direction int // 0 - UP, 1 - RIGHT, 2 - DOWN, 3 - LEFT
}

func createMap(rows []string) MapState {
	var state MapState

	state.height = len(rows)
	state.width = len(rows[0])
	state.grid = make([][]rune, state.height)

	for i, row := range rows {
		state.grid[i] = make([]rune, state.width)
		for j, char := range row {
			if char == '^' {
				state.guard = [2]int{i, j}
			}
			state.grid[i][j] = char
		}
	}

	return state
}

func getNextGuardPosition(state *MapState) (next [2]int, outOfBounds bool) {
	// Get the next position of the guard
	// based on the current state

	// Calculate the diff based on direction
	dy := 0
	dx := 0

	switch state.direction {
	case 0:
		dy = -1
	case 1:
		dx = 1
	case 2:
		dy = 1
	case 3:
		dx = -1
	}

	// Add the diff to get next pos, and check if it's OOB
	nextPos := [2]int{state.guard[0] + dy, state.guard[1] + dx}
	if nextPos[0] < 0 || nextPos[0] >= state.height || nextPos[1] < 0 || nextPos[1] >= state.width {
		return nextPos, true
	}

	// Check if the next pos is a wall
	// If it is, it stays in the same place
	if state.grid[nextPos[0]][nextPos[1]] == '#' {
		return state.guard, false
	}

	// No wall, not oob, we can move
	return nextPos, false
}

func getGuardPathset(state MapState) map[[2]int]bool {
	// Create a set of all the guard's path points
	pathset := make(map[[2]int]bool)

	nextPos, oob := state.guard, false
	for !oob {
		pathset[nextPos] = true
		nextPos, oob = getNextGuardPosition(&state)

		// If we hit a wall, we'll have the same nextPos so we change the direction
		if nextPos == state.guard {
			state.direction = (state.direction + 1) % 4
		} else {
			state.guard = nextPos
		}
	}

	return pathset
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
