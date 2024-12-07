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

	// In the data string, find all matches for all the commands
	commandsPattern := regexp.MustCompile(`do\(\)|mul\(\d{1,3},\d{1,3}\)|don't\(\)`)
	commands := commandsPattern.FindAllString(dataStr, -1)

	// Iteratively use mul based on do status
	doStatus := true
	total := 0
	for _, command := range commands {
		switch command {
		case "do()":
			doStatus = true
		case "don't()":
			doStatus = false
		default:
			if !doStatus {
				continue
			}
			fmt.Println(command)
			nums := extractNumbers(command)
			total += nums[0] * nums[1]
		}
	}

	// We're done!
	fmt.Println(total)
}

func extractNumbers(match string) []int {

	// Extract the two numbers from the match
	pattern := regexp.MustCompile(`([0-9]{1,3})`)
	numbers := pattern.FindAllString(match, -1)

	if len(numbers) != 2 {
		panic("Invalid match")
	}

	return []int{toInt(numbers[0]), toInt(numbers[1])}
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
