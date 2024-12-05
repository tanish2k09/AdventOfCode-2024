package main

import (
	"fmt"
	"os"
	"strconv"
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

	// We don't know how many numbers are per record, we'll just do it iteratively
	records := strings.Split(dataStr, "\n")

	// For each record, check the safety of it and update the total safe record count
	safe := 0
	for _, recordStr := range records {
		// Convert the record to ints
		record := toIntRecord(strings.Split(recordStr, " "))
		if isSafe(record) {
			safe++
		}
	}

	// We're done!
	fmt.Println(safe)
}

func toIntRecord(record []string) []int {
	intRecord := make([]int, len(record))
	for i, numStr := range record {
		intRecord[i] = toInt(numStr)
	}
	return intRecord
}

func isSafe(record []int) bool {
	// Each record is either allowed to be continuously increasing or continuously decreasing
	// Each consecutive value must differ by at least 1 and at most 3 to be safe.
	if len(record) < 2 { // Base condition so we can find the direction
		return true
	}

	direction := sign(record[1] - record[0])
	if direction == 0 {
		return false
	}

	for i := 1; i < len(record); i++ {
		diff := record[i] - record[i-1]
		if sign(diff) != direction || abs(diff) > 3 || abs(diff) < 1 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func toInt(x string) int {
	i, err := strconv.Atoi(x)
	check(err)
	return i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
