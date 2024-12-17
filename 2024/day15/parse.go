package main

import (
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func Map(in string) (matrix [][]byte) {
	_ = utils.ScanFileLineByLine(strings.NewReader(in), func(line string) {
		matrix = append(matrix, []byte(line))
	})
	return
}

func Movements(in string) (moves []byte) {
	_ = utils.ScanFileLineByLine(strings.NewReader(in), func(line string) {
		moves = append(moves, []byte(line)...)
	})
	return
}
