package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

// Rules
// - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// - If the stone is engraved with a number that has an even number of digits,
//   it is replaced by two stones. The left half of the digits are engraved on the new left stone,
//   and the right half of the digits are engraved on the new right stone.
//   (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// - If none of the other rules apply, the stone is replaced by a new stone;
//   the old stone's number multiplied by 2024 is engraved on the new stone.

// No matter how the stones change, their order is preserved, and they stay on their perfectly straight line.

var singleStoneCache = make(map[int][]int)

func ApplyRules(stone int) []int {
	// look for already calculated value
	if res, ok := singleStoneCache[stone]; ok {
		return res
	}

	// rule 1
	if stone == 0 {
		res := []int{1}
		singleStoneCache[stone] = res
		return res
	}

	// rule 2
	// how many times you can divide the number by 10 before it becomes less than 1
	// every time we divide by 10 we "loose" one digit
	numDigits := int(math.Log10(float64(stone))) + 1
	if numDigits%2 == 0 {
		// creates a power of 10 that matches half the length of the number, example 5522: 4 digits, 2**10 = 100
		divisor := int(math.Pow10(numDigits / 2))
		left := stone / divisor  // example 5522: 5522/100 = 55,22
		right := stone % divisor // example 5522: 5522%100 = 22
		res := []int{left, right}
		singleStoneCache[stone] = res
		return res
	}

	// rule 3
	res := []int{stone * 2024}
	singleStoneCache[stone] = res

	return res
}

type countKey [2]int

var countCache = make(map[countKey]int)

func CountByStone(stone int, times int) int {
	cacheKey := countKey{stone, times}
	if value, found := countCache[cacheKey]; found {
		return value
	}

	// make a calculation of a rules
	newStones := ApplyRules(stone)

	// the basic case, the last calculation step, simply return how many we've got
	if times == 1 {
		return len(newStones)
	}

	// for each of the stones, do the calc again but 1 time less
	var count int
	for _, newStone := range newStones {
		count += CountByStone(newStone, times-1)
	}

	countCache[cacheKey] = count

	return count
}

func BlinkTimes(lineOfStones string, times int) int {
	// create ints for faster calculations
	var stones []int
	for _, s := range strings.Split(lineOfStones, " ") {
		stones = append(stones, utils.Int(s))
	}

	var count int
	for _, stone := range stones {
		count += CountByStone(stone, times)
	}
	return count
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	numOfStones := BlinkTimes(exampleInput, 75)
	fmt.Printf("Number of stones after blink: %d\n", numOfStones)
}
