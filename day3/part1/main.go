package main

import (
	"fmt"
	"os"
	"regexp"
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

	// In the data string, find all matches for the mul regex
	pattern := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))`)
	matches := pattern.FindAllString(dataStr, -1)

	// Now we extract the numbers from each match
	pairs := extractNumbers(matches)

	// Now we multiply the pairs and sum them up
	total := getPairMultipliedSum(pairs)

	// We're done!
	fmt.Println(total)
}

func extractNumbers(matches []string) [][]int {
	list := make([][]int, len(matches))

	for index, match := range matches {
		// Extract the two numbers from the match
		pattern := regexp.MustCompile(`([0-9]{1,3})`)
		numbers := pattern.FindAllString(match, -1)

		if len(numbers) != 2 {
			panic("Invalid match")
		}

		list[index] = []int{toInt(numbers[0]), toInt(numbers[1])}
	}

	return list
}

func getPairMultipliedSum(pairs [][]int) int {
	sum := 0
	for _, pair := range pairs {
		sum += pair[0] * pair[1]
	}
	return sum
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
