package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func CountXXXMAS(in [][]byte) int {
	word := []byte("MAS")
	var total int

	for y, row := range in {
		for x := range row {
			// MAS can be written forwards or backwards
			// from the top or bottom, left or right
			// which boils down to: if M is on the top left then bottom right could only be S
			// that means on all 4 corners can be only M or S and in total only 2 of each

			// update:
			// after some time i understood that MAM and SAS also fall into that trap ;)

			// starting from the central A
			if in[y][x] != 'A' {
				continue
			}

			var masCount int

			// top-left
			if wordFromPosition(in, word, y-1, x-1, utils.SouthEast) {
				masCount++
			}
			// top-right
			if wordFromPosition(in, word, y-1, x+1, utils.SouthWest) {
				masCount++
			}
			// bottom-left
			if wordFromPosition(in, word, y+1, x-1, utils.NorthEast) {
				masCount++
			}
			// bottom-right
			if wordFromPosition(in, word, y+1, x+1, utils.NorthWest) {
				masCount++
			}

			if masCount == 2 {
				total++
			}
		}
	}

	return total
}

func CountXMAS(in [][]byte) int {
	word := []byte("XMAS")
	var total int

	for y, row := range in {
		for x := range row {
			for _, dir := range utils.AllDirections {
				if wordFromPosition(in, word, y, x, dir) {
					total++
				}
			}
		}
	}

	return total
}

func wordFromPosition(in [][]byte, word []byte, y, x int, dir utils.Direction) bool {
	// check boundaries of starting position
	if y < 0 || y >= len(in) {
		return false
	}
	if x < 0 || x >= len(in[y]) {
		return false
	}

	for step, ch := range word {
		// calculate coordinated of the character beginning + number * movement
		nextY := y + step*dir.Y()
		nextX := x + step*dir.X()

		// check row boundaries
		if nextY < 0 || nextY >= len(in) {
			return false
		}

		// check column boundaries
		if nextX < 0 || nextX >= len(in[nextY]) {
			return false
		}

		// not expected character
		if in[nextY][nextX] != ch {
			return false
		}
	}

	return true
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	data, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(exampleInput))
	if err != nil {
		panic(err)
	}

	countXMAS := CountXMAS(data)
	fmt.Printf("Total XMAS words: %d\n", countXMAS)

	countXXXMAS := CountXXXMAS(data)
	fmt.Printf("Total X-MAS words: %d\n", countXXXMAS)
}
