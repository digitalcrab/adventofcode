package main

import (
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func Parse(in string) (pos []utils.Pos) {
	_ = utils.ScanFileLineByLine(strings.NewReader(in), func(line string) {
		ss := utils.StringsToInts(strings.Split(line, ","))
		pos = append(pos, utils.NewPos(ss[1], ss[0]))
	})
	return
}
