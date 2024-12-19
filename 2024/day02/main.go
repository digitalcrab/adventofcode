package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

// IsReportSafe calculates if report is safe or not by the following rules
// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func IsReportSafe(levels []int) bool {
	// calculate initial direction
	increasing := levels[1] > levels[0]

	// start from the second element
	for i := 1; i < len(levels); i++ {
		// calculate absolute diff
		diff := int(math.Abs(float64(levels[i] - levels[i-1])))

		// check if difference within the range
		if diff < 1 || diff > 3 {
			return false
		}

		if levels[i] > levels[i-1] && !increasing {
			// if it's not increasing but was not
			return false
		} else if levels[i] < levels[i-1] && increasing {
			// and other way around
			return false
		}
	}

	return true
}

// IsReportSafeWithTolerance calculates if levels are safe.
// Now, the same rules apply as before, except if removing a single level from an unsafe report
// would make it safe, the report instead counts as safe.
func IsReportSafeWithTolerance(levels []int) bool {
	// first checking without any modifications
	if IsReportSafe(levels) {
		return true
	}

	// lets try removing level by level and see if something changes
	// TODO: not optimal ;)
	for i := range levels {
		// copy first portion to the empty slice
		newLevels := append([]int{}, levels[:i]...)
		// copy the rest of the slice skipping the current element
		newLevels = append(newLevels, levels[i+1:]...)

		isSafe := IsReportSafe(newLevels)
		if isSafe {
			return true
		}
	}

	return false
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func NumberOfSafeReports(input string, debug bool) (int, int) {
	var totalSafe, totalSafeTolerance int
	err := utils.ScanFileLineByLine(strings.NewReader(input), func(line string) {
		levels := utils.StringsToInts(strings.Split(line, " "))
		isSafe := IsReportSafe(levels)
		if isSafe {
			totalSafe++
		}
		isSafeTolerance := IsReportSafeWithTolerance(levels)
		if isSafeTolerance {
			totalSafeTolerance++
		}
		if debug {
			fmt.Printf("report %s = %t (%t)\n", line, isSafe, isSafeTolerance)
		}
	})
	if err != nil {
		panic(err)
	}
	return totalSafe, totalSafeTolerance
}

func main() {
	totalSafe, totalSafeTolerance := NumberOfSafeReports(exampleInput, true)
	fmt.Printf("Total safe reports: %d (%d)\n", totalSafe, totalSafeTolerance)
}
