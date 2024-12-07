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

	// Get the looping wall test count
	combos := testGuardPathset(mapState, pathset)

	// We're done!
	fmt.Println(combos)
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

func (ms *MapState) Clone() *MapState {
	// Deep copy the grid (2D slice of runes)
	gridCopy := make([][]rune, len(ms.grid))
	for i := range ms.grid {
		gridCopy[i] = make([]rune, len(ms.grid[i]))
		copy(gridCopy[i], ms.grid[i])
	}

	// Return a new instance of MapState
	return &MapState{
		width:     ms.width,
		height:    ms.height,
		grid:      gridCopy,
		guard:     ms.guard,
		direction: ms.direction,
	}
}

func setRuneAt(state MapState, pos [2]int, character rune) (oldChar rune) {
	// Create a new map state with a wall at the given position
	oldChar = state.grid[pos[0]][pos[1]]
	state.grid[pos[0]][pos[1]] = character
	return oldChar
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

func testGuardPathset(state MapState, pathset map[[2]int]bool) int {
	total := 0
	startingGuard := state.guard

	// We're gonna add walls at each point the guard takes, except his starting point
	delete(pathset, startingGuard)

	for wallPos := range pathset {
		// Create a clone map state
		testState := *state.Clone()

		// Set the guard to the starting position
		testState.guard = startingGuard
		testState.direction = 0

		// Create a wall at the wall position
		setRuneAt(testState, wallPos, '#')

		// Check if guard escapes this new map, if they don't, add the total
		if !doesGuardEscape(testState) {
			total++
		}
	}

	return total
}

func doesGuardEscape(state MapState) bool {
	// Create a set of all the guard's path points
	pathset := make(map[[3]int]bool)

	tracker := [3]int{state.guard[0], state.guard[1], state.direction} // Tracks direction as well
	pathset[tracker] = true                                            // Mark guard's current position

	oob := false
	for !oob {
		var nextPos [2]int
		nextPos, oob = getNextGuardPosition(&state)
		tracker = [3]int{nextPos[0], nextPos[1], state.direction}

		// If we've already seen the next position before, we're in a loop
		if nextPos == state.guard {
			state.direction = (state.direction + 1) % 4
		} else if !pathset[tracker] {
			state.guard = nextPos
		} else {
			return false
		}

		pathset[tracker] = true
	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
