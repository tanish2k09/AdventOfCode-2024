package main

import (
	"fmt"
	"os"
	"sort"
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

	// Sort the two columns
	sort.Ints(left)
	sort.Ints(right)

	// Calculate distances
	distances := getDistances(left, right)

	// Calculate the sum of the distances
	totalDistance := sum(distances)

	// We're done!
	fmt.Println(totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDistances(left []int, right []int) []int {
	distances := make([]int, len(left))
	for i := 0; i < len(left); i++ {
		distances[i] = abs(right[i] - left[i])
	}
	return distances
}

func sum(x []int) int {
	total := 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	return total
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
