package main

import (
	"container/heap"
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/digitalcrab/adventofcode/utils"
)

const (
	StateEmpty     = iota
	StateCorrupted // basically here is the byte longing
	StatePath

	mapHeight = 71
	mapWidth  = 71
)

var start, end = utils.NewPos(0, 0), utils.NewPos(mapHeight-1, mapWidth-1)

func FindShortestPath(obstacles map[utils.Pos]int) []utils.Pos {
	// score storage for all points, start with 0 for the beginning and rest amx out
	scores := make(map[utils.Pos]int, mapHeight*mapWidth)
	for p := range utils.PositionsForHeightWidth(mapHeight, mapWidth) {
		scores[p] = 1<<32 - 1 // max int without minus
	}
	scores[start] = 0

	// store how to get to the cell
	predecessors := make(map[utils.Pos]utils.Pos)

	// create a queue of steps
	queue := &utils.Queue{}
	heap.Init(queue)
	heap.Push(queue, &utils.Step{
		St: utils.StepState{P: start}, // direction we do not care much here
		S:  0,
	})

	for queue.Len() > 0 {
		current := queue.Pop().(*utils.Step)

		// from here we have 4 options: go in any direction if this is not the end of the map
		for _, dir := range utils.AzimuthDirections {
			nextPos := current.St.P.Next(dir)

			// check boundaries and corrupted cells
			if !nextPos.InBoundaries(mapHeight, mapWidth) || obstacles[nextPos] != StateEmpty {
				continue
			}

			// if a new score is better (lower) then we go this way
			if nextScore := scores[current.St.P] + 1; nextScore < scores[nextPos] {
				// update score
				scores[nextPos] = nextScore
				// explore next step
				heap.Push(queue, &utils.Step{
					St: utils.StepState{P: nextPos},
					S:  nextScore,
				})
				// save how we get here (mostly for visualisation, you actually do not need it to answer )
				predecessors[nextPos] = current.St.P
			}
		}
	}

	// no path found
	if scores[end] == 1<<32-1 {
		return nil
	}

	// backtrack all nodes
	var nodes = []utils.Pos{end}
	currentPos := end

	for currentPos != start {
		from := predecessors[currentPos]
		nodes = append(nodes, from)
		currentPos = from
	}

	return nodes
}

func FindPathWithObstacles(fallingBytes []utils.Pos) []utils.Pos {
	obstacles := make(map[utils.Pos]int)
	for _, bp := range fallingBytes {
		obstacles[bp] = StateCorrupted
	}
	return FindShortestPath(obstacles)
}

func BinaryFind(left, right int, fallingBytes []utils.Pos) utils.Pos {
	if left >= right {
		return fallingBytes[left]
	}

	// get the middle point
	mid := left + (right-left)/2

	// run search for left part
	path := FindPathWithObstacles(fallingBytes[:mid+1])

	// if there is no path already then it must be from the left to middle point
	if len(path) == 0 {
		return BinaryFind(left, mid, fallingBytes)
	}

	// if there is a way to get to the end then it's gonna be a way till mid point
	return BinaryFind(mid+1, right, fallingBytes)
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	fallingBytes := Parse(DayInput)

	// find the path with 1024 bytes
	path := FindPathWithObstacles(fallingBytes[:1024])
	fmt.Printf("[Part 1]: %d\n", len(path)-1)

	// looping through every byte not feasible, maybe we can use binary search to reduce the problem size?
	// up until 1024 we already know from the part 1
	breakThroughPos := BinaryFind(1024, len(fallingBytes)-1, fallingBytes)
	fmt.Printf("[Part 2]: %v\n", breakThroughPos)

	vis := NewVisualisation(fallingBytes)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("RAM Run")
	ebiten.SetWindowSize(1024, 1024)
	if err := ebiten.RunGame(vis); err != nil {
		panic(err)
	}
}
