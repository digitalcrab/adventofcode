package main

import (
	_ "embed"
	"fmt"
	"regexp"

	"github.com/digitalcrab/adventofcode/utils"
)

var (
	re = regexp.MustCompile(`(?P<statement>mul\((?P<num1>\d{1,3}),(?P<num2>\d{1,3})\)|do\(\)|don't\(\))`)
)

func DecodeMemory(data string, doLogic bool) int {
	// initially all should be applied
	data = "do()" + data

	var sum int
	var apply bool

loop:
	for _, matches := range re.FindAllStringSubmatch(data, -1) {

		// the rule to apply or not
		switch matches[1] { // statement
		case "do()":
			apply = true
			continue loop
		case "don't()":
			apply = false
			continue loop
		}

		// the next statement we should skip
		if doLogic && !apply {
			continue
		}

		num1 := utils.Int(matches[2]) // num1
		num2 := utils.Int(matches[3]) // num2
		sum += num1 * num2
	}
	return sum
}

//go:embed "example.txt"
var exampleInput string

func main() {
	sum := DecodeMemory(exampleInput, true)
	fmt.Printf("Memory calculation: %d\n", sum)
}
