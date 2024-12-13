package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/digitalcrab/adventofcode/utils"
)

func SumOfTrailheadScores(in [][]byte, uniqueWays bool) (sum int) {
	for y, row := range in {
		for x := range row {
			// start trail only from 0
			if in[y][x] != '0' {
				continue
			}
			var seen map[utils.Pos]struct{}
			if uniqueWays {
				seen = make(map[utils.Pos]struct{})
			}
			// add scores together
			sum += DFS(y, x, in, seen)
		}
	}

	return
}

func DFS(y, x int, in [][]byte, seen map[utils.Pos]struct{}) (res int) {
	if seen != nil {
		pos := utils.NewPos(y, x)
		// we already went this way before so this route does not count
		if _, found := seen[pos]; found {
			return 0
		}
		// remember that we've visited this point
		seen[pos] = struct{}{}
	}

	// check if we've reached the top point, if so, we add 1 as the result
	if in[y][x] == '9' {
		return 1
	}

	// moving only in 4 directions (North, East, South, West)
	for _, dir := range utils.AzimuthDirections {
		// the next step
		newY, newX := y+dir.Y, x+dir.X
		// check boundaries
		if newY < 0 || newY >= len(in) || newX < 0 || newX >= len(in[newY]) {
			continue
		}
		// we can move only if next step is 1 more then previous
		if in[y][x]+1 != in[newY][newX] {
			continue
		}
		// go to the next step
		res += DFS(newY, newX, in, seen)
	}

	return
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	data, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(DayInput))
	if err != nil {
		panic(err)
	}

	sum := SumOfTrailheadScores(data, true)
	fmt.Printf("Sum of all trainhead scores: %d\n", sum)

	vis := NewVisualisation(data)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("Hoof It")
	ebiten.SetWindowSize(1024, 1024)
	if err := ebiten.RunGame(vis); err != nil {
		panic(err)
	}
}
