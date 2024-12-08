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
	rows := strings.Split(dataStr, "\n")

	// Parse the rows into expressions
	expressions := parseRows(rows)

	// Test each expression
	count := 0
	for _, e := range expressions {
		testExpr := testExpr(&e, e.factors[0], 1)
		if testExpr {
			count += e.target
		}
	}

	// We're done!
	fmt.Println(count)
}

type Expr struct {
	target  int
	factors []int
}

func parseRows(rows []string) []Expr {
	exprs := make([]Expr, len(rows))
	numPattern := regexp.MustCompile(`\d+`)

	for k, v := range rows {
		nums := numPattern.FindAllString(v, -1)
		target := toInt(nums[0])
		factors := make([]int, len(nums)-1)
		for i, n := range nums[1:] {
			factors[i] = toInt(n)
		}
		exprs[k] = Expr{target, factors}
	}
	return exprs
}

func testExpr(e *Expr, total int, next int) bool {
	// FIX: Out of range
	if next >= len(e.factors) {
		return total == e.target
	}

	// Operated values
	added := total + e.factors[next]
	mult := total * e.factors[next]

	// If next is the last element, test both operators
	if next == len(e.factors) {
		if added == e.target || mult == e.target {
			return true
		}
	}

	// Test addition, call recursively and return early if true is received
	plusRoute := testExpr(e, added, next+1)
	if plusRoute {
		return true
	}

	// Test multiplication
	multRoute := testExpr(e, mult, next+1)
	return multRoute
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
