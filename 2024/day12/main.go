package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func TotalPrice(in [][]byte) (sumPerimeter, sumSides int) {
	seen := make(map[utils.Pos]struct{})

	for y, row := range in {
		for x := range row {
			pos := utils.NewPos(y, x)
			if _, found := seen[pos]; found {
				continue
			}

			sides := make(map[string][]int)
			// perimeter from that function used in part 1, but we also have the
			// same value after completing part to and int's in `bricksCount`
			area, _ := DFS(y, x, in, seen, sides)

			wallsCount, bricksCount := calcSides(sides)
			sumPerimeter += area * bricksCount
			sumSides += area * wallsCount
		}
	}

	return
}

func calcSides(sides map[string][]int) (wallsCount, bricksCount int) {
	// key is unique wal and values are uninterrupted bricks
	for _, bricks := range sides {
		bricksCount += len(bricks)

		if len(bricks) == 1 {
			wallsCount++
			continue
		}

		// sorting is necessary as we move along the wall in the direction of growth
		// we should be able to see gaps
		sort.Ints(bricks)

		// calc how many times brick interrupted
		interruptions := 0
		for i := 1; i < len(bricks); i++ {
			// if diff between them more than 1 brick
			if bricks[i]-bricks[i-1] > 1 {
				interruptions++
			}
		}

		wallsCount += interruptions + 1 // number of gaps + one initial wall
	}

	return
}

func DFS(y, x int, in [][]byte, seen map[utils.Pos]struct{}, sides map[string][]int) (area, perimeter int) {
	pos := utils.NewPos(y, x)
	// we already went this way before so this route does not count
	if _, found := seen[pos]; found {
		return 0, 0
	}

	// remember that we've visited this point
	seen[pos] = struct{}{}
	area++

	// moving only in 4 directions (North, East, South, West)
	for _, dir := range utils.AzimuthDirections {
		// the next step
		newY, newX := y+dir.Y(), x+dir.X()

		// check boundaries
		if newY < 0 || newY >= len(in) || newX < 0 || newX >= len(in[newY]) {
			// out boundaries means need fence
			perimeter++
			recordSide(y, x, dir, sides)
			continue
		}

		// move only within the same region
		if in[y][x] != in[newY][newX] {
			// not the same region? needs fence
			perimeter++
			recordSide(y, x, dir, sides)
			continue
		}

		// go to the next step
		newArea, newPerimeter := DFS(newY, newX, in, seen, sides)
		area += newArea
		perimeter += newPerimeter
	}

	return
}

func recordSide(y, x int, dir utils.Direction, sides map[string][]int) {
	// key is going to represent the wall that we are moving round
	var key string
	// value represent individual bricks (basically perimeter or that wall)
	var value int

	switch {
	case dir.Y() != 0: // we moved caning the row, means hit the wall here, then use column as a brick
		key = fmt.Sprintf("%v:%d", dir, y)
		value = x
	case dir.X() != 0:
		key = fmt.Sprintf("%v:%d", dir, x)
		value = y
	}

	// we collect wall and bricks, we also included direction into the wall name
	sides[key] = append(sides[key], value)
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	data, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(DayInput)) // 1304712 to low
	if err != nil {
		panic(err)
	}

	sumPerimeter, sumSides := TotalPrice(data)
	fmt.Printf("Total price %d for perimeter based or discount %d\n", sumPerimeter, sumSides)
}
