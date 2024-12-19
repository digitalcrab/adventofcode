package main

import (
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func Parse(in string) (Machine, []int) {
	var lines []string
	_ = utils.ScanFileLineByLine(strings.NewReader(in), func(line string) {
		if line == "" {
			return
		}
		lines = append(lines, line)
	})

	rl := len("Register A: ")
	pl := len("Program: ")

	machine := Machine{
		A: utils.Int(lines[0][rl:]),
		B: utils.Int(lines[1][rl:]),
		C: utils.Int(lines[2][rl:]),
	}

	return machine, utils.StringsToInts(strings.Split(lines[3][pl:], ","))
}
