package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

func TotalPrice(in [][]byte) (sumPerimeter, sumSides int) {
	seen := make(map[utils.Position]struct{})

	for rx, row := range in {
		for cx := range row {
			pos := utils.NewPosition(rx, cx)
			if _, found := seen[pos]; found {
				continue
			}

			sides := make(map[string][]int)
			// perimeter from that function used in part 1, but we also have the
			// same value after completing part to and int's in `bricksCount`
			area, _ := DFS(rx, cx, in, seen, sides)

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

func DFS(rx, cx int, in [][]byte, seen map[utils.Position]struct{}, sides map[string][]int) (area, perimeter int) {
	pos := utils.NewPosition(rx, cx)
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
		newRow, newCol := rx+dir.Row, cx+dir.Col

		// check boundaries
		if newRow < 0 || newRow >= len(in) || newCol < 0 || newCol >= len(in[newRow]) {
			// out boundaries means need fence
			perimeter++
			recordSide(rx, cx, dir, in, sides)
			continue
		}

		// move only within the same region
		if in[rx][cx] != in[newRow][newCol] {
			// not the same region? needs fence
			perimeter++
			recordSide(rx, cx, dir, in, sides)
			continue
		}

		// go to the next step
		newArea, newPerimeter := DFS(newRow, newCol, in, seen, sides)
		area += newArea
		perimeter += newPerimeter
	}

	return
}

func recordSide(rx, cx int, dir *utils.Direction, in [][]byte, sides map[string][]int) {
	// key is going to represent the wall that we are moving round
	var key string
	// value represent individual bricks (basically perimeter or that wall)
	var value int

	newRow, newCol := rx+dir.Row, cx+dir.Col

	switch {
	case newRow < 0 || newRow >= len(in):
		key = fmt.Sprintf("%s:row:%d", dir, newRow)
		value = newCol
	case newCol < 0 || newCol >= len(in[newRow]):
		key = fmt.Sprintf("%s:col:%d", dir, newCol)
		value = newRow
	case in[rx][cx] != in[newRow][newCol] && dir.Row != 0:
		key = fmt.Sprintf("%s:row:%d", dir, newRow)
		value = newCol
	case in[rx][cx] != in[newRow][newCol] && dir.Col != 0:
		key = fmt.Sprintf("%s:col:%d", dir, newCol)
		value = newRow
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
