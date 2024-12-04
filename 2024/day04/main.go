package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

type direction struct {
	row, col int
}

var (
	dN  = direction{row: -1, col: 0}
	dNE = direction{row: -1, col: 1}
	dE  = direction{row: 0, col: 1}
	dSE = direction{row: 1, col: 1}
	dS  = direction{row: 1, col: 0}
	dSW = direction{row: 1, col: -1}
	dW  = direction{row: 0, col: -1}
	dNW = direction{row: -1, col: -1}

	allDirections = []direction{dN, dNE, dE, dSE, dS, dSW, dW, dNW}
)

func CountXXXMAS(in [][]byte) int {
	word := []byte("MAS")
	var total int

	for rx, row := range in {
		for cx := range row {
			// MAS can be written forwards or backwards
			// from the top or bottom, left or right
			// which boils down to: if M is on the top left then bottom right could only be S
			// that means on all 4 corners can be only M or S and in total only 2 of each

			// update:
			// after some time i understood that MAM and SAS also fall into that trap ;)

			// starting from the central A
			if in[rx][cx] != 'A' {
				continue
			}

			var masCount int

			// top-left
			if wordFromPosition(in, word, rx-1, cx-1, dSE) {
				masCount++
			}
			// top-right
			if wordFromPosition(in, word, rx-1, cx+1, dSW) {
				masCount++
			}
			// bottom-left
			if wordFromPosition(in, word, rx+1, cx-1, dNE) {
				masCount++
			}
			// bottom-right
			if wordFromPosition(in, word, rx+1, cx+1, dNW) {
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

	for rx, row := range in {
		for cx := range row {
			for _, dir := range allDirections {
				if wordFromPosition(in, word, rx, cx, dir) {
					total++
				}
			}
		}
	}

	return total
}

func wordFromPosition(in [][]byte, word []byte, row, col int, dir direction) bool {
	// check boundaries of starting position
	if row < 0 || row >= len(in) {
		return false
	}
	if col < 0 || col >= len(in[row]) {
		return false
	}

	for step, ch := range word {
		// calculate coordinated of the character beginning + number * movement
		chRow := row + step*dir.row
		chCol := col + step*dir.col

		// check row boundaries
		if chRow < 0 || chRow >= len(in) {
			return false
		}

		// check column boundaries
		if chCol < 0 || chCol >= len(in[chRow]) {
			return false
		}

		// not expected character
		if in[chRow][chCol] != ch {
			return false
		}
	}

	return true
}

//go:embed "example.txt"
var exampleInput string

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
