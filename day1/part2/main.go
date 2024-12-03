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

	// The data is formatted as "x   y\n" so we build two arrays from the columns
	dataSplit := strings.Split(dataStr, "\n")
	left := make([]int, len(dataSplit))
	right := make([]int, len(dataSplit))
	for k, v := range dataSplit {
		nums := strings.Split(v, "   ")
		left[k] = toInt(nums[0])
		right[k] = toInt(nums[1])
	}

	// Create a counter map from the right column
	counter := count(right)

	// Calculate similarity score with left values and counts
	sim := similarity(left, counter)

	// We're done!
	fmt.Println(sim)
}

func similarity(left []int, counter map[int]int) int {
	total := 0

	for _, v := range left {
		total += v * counter[v]
	}

	return total
}

func count(data []int) map[int]int {
	counter := make(map[int]int)

	// Count the number of times each number appears in the data
	for _, v := range data {
		if current, present := counter[v]; present {
			counter[v] = current + 1
		} else {
			counter[v] = 1
		}
	}

	return counter
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
