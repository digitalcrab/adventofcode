package main

import (
	"regexp"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

var rx = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func Parse(in string) (robots []*Robot) {
	_ = utils.ScanFileLineByLine(strings.NewReader(in), func(line string) {
		matches := rx.FindStringSubmatch(line)
		robots = append(robots, &Robot{
			pos:      utils.NewPos(utils.Int(matches[2]), utils.Int(matches[1])),
			velocity: utils.NewPos(utils.Int(matches[4]), utils.Int(matches[3])),
		})
	})
	return
}

func Display(robots []*Robot, height, weight int) {
	matrix := utils.CreateMatrix(height, weight, '.')
	for _, r := range robots {
		if matrix[r.pos.Y()][r.pos.X()] == '.' {
			matrix[r.pos.Y()][r.pos.X()] = '0'
		}
		nextValue := matrix[r.pos.Y()][r.pos.X()] - '0'
		nextValue++
		matrix[r.pos.Y()][r.pos.X()] = nextValue + '0'
	}
	utils.PrintMatrix(matrix)
}
