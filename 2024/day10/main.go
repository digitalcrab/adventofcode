package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func SumOfTrailheadScores(in [][]byte, uniqueWays bool) (sum int) {
	for rx, row := range in {
		for cx := range row {
			// start trail only from 0
			if in[rx][cx] != '0' {
				continue
			}
			var seen map[utils.Position]struct{}
			if uniqueWays {
				seen = make(map[utils.Position]struct{})
			}
			// add scores together
			sum += DFS(rx, cx, in, seen)
		}
	}

	return
}

func DFS(rx, cx int, in [][]byte, seen map[utils.Position]struct{}) (res int) {
	if seen != nil {
		pos := utils.NewPosition(rx, cx)
		// we already went this way before so this route does not count
		if _, found := seen[pos]; found {
			return 0
		}
		// remember that we've visited this point
		seen[pos] = struct{}{}
	}

	// check if we've reached the top point, if so, we add 1 as the result
	if in[rx][cx] == '9' {
		return 1
	}

	// moving only in 4 directions (North, East, South, West)
	for _, dir := range utils.AzimuthDirections {
		// the next step
		newRow, newCol := rx+dir.Row, cx+dir.Col
		// check boundaries
		if newRow < 0 || newRow >= len(in) || newCol < 0 || newCol >= len(in[newRow]) {
			continue
		}
		// we can move only if next step is 1 more then previous
		if in[rx][cx]+1 != in[newRow][newCol] {
			continue
		}
		// go to the next step
		res += DFS(newRow, newCol, in, seen)
	}

	return
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

	sum := SumOfTrailheadScores(data, true)
	fmt.Printf("Sum of all trainhead scores: %d\n", sum)

}
