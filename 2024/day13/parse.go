package main

import (
	"regexp"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

var (
	rxButton = regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)`)
	rxPrice  = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
)

func Parse(data string) (machines []Machine) {
	var current *Machine
	_ = utils.ScanFileLineByLine(strings.NewReader(data), func(line string) {
		if line == "" {
			machines = append(machines, *current)
			current = nil
			return
		}

		if current == nil {
			current = &Machine{}
		}

		if rxButton.MatchString(line) {
			var btn Button
			if matches := rxButton.FindStringSubmatch(line); matches[1] == "A" {
				btn = Button{price: 3, move: utils.NewPos(utils.Int(matches[3]), utils.Int(matches[2]))}
			} else {
				btn = Button{price: 1, move: utils.NewPos(utils.Int(matches[3]), utils.Int(matches[2]))}
			}
			current.buttons = append(current.buttons, btn)
			return
		}

		if rxPrice.MatchString(line) {
			matches := rxPrice.FindStringSubmatch(line)
			current.price = utils.NewPos(utils.Int(matches[2]), utils.Int(matches[1]))
		}
	})

	if current != nil {
		machines = append(machines, *current)
	}

	return
}
