package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

type Guardian struct {
	pos    utils.Pos
	dirIdx int
}

func (g *Guardian) Rotate() {
	g.dirIdx = (g.dirIdx + 1) % len(utils.AzimuthDirections)
}

func FindGuardian(in [][]byte) Guardian {
	for p, ch := range utils.PositionsValues(in) {
		if ch != '.' && ch != '#' {
			guardian := Guardian{
				pos:    p,
				dirIdx: utils.SymbolDirectionIdx[ch],
			}
			return guardian
		}
	}
	panic("not found")
}

// Walk walks guardian to the exit of the map and returns number of distinct steps he's made
// Very strict patrol protocol
// If there is something directly in front of you, turn right 90 degrees.
// Otherwise, take a step forward.
// How many distinct positions will the guard visit before leaving the mapped area?
func Walk(in [][]byte, guardian Guardian) (int, int) {
	// store distinkt steps in the map where key is a combination of
	// row and col, value does not matter actually for now
	visitedDirections := map[utils.Pos]map[int]struct{}{}
	var possibleObstacle int

	looped := walk(in, guardian, visitedDirections)

	// something terrible
	if looped {
		panic("should not loop for the first time")
	}

	distinctSteps := len(visitedDirections)
	loopsCh := make(chan bool)

	// loop over all positions guardian went through
	for pos := range visitedDirections {
		// skip initial post
		if pos.Eq(guardian.pos) {
			continue
		}

		go func() {
			// place a new obstacle
			copyIn := utils.DuplicateBytesMatrix(in)
			copyIn[pos.Y()][pos.X()] = '#'

			// run again the thing
			loopsCh <- walk(copyIn, guardian, map[utils.Pos]map[int]struct{}{})
		}()
	}

	for pos := range visitedDirections {
		if pos.Eq(guardian.pos) {
			continue
		}
		if <-loopsCh {
			possibleObstacle++
		}
	}

	return distinctSteps, possibleObstacle
}

func walk(in [][]byte, guardian Guardian, visitedDirections map[utils.Pos]map[int]struct{}) bool {
	for {
		// record current position + direction
		if _, visited := visitedDirections[guardian.pos]; !visited {
			visitedDirections[guardian.pos] = map[int]struct{}{guardian.dirIdx: {}}
		} else {
			// if visited, we need to detect loop
			if _, looped := visitedDirections[guardian.pos][guardian.dirIdx]; looped {
				return true
			}
			// not looped then just save
			visitedDirections[guardian.pos][guardian.dirIdx] = struct{}{}
		}

		// calculate next step as a current step plus directional change
		dir := utils.AzimuthDirections[guardian.dirIdx]
		nextPos := guardian.pos.Next(dir)
		nextY, nextX := nextPos.Values()

		// check boundaries
		if nextY < 0 || nextY >= len(in) || nextX < 0 || nextX >= len(in[nextY]) {
			// guardian leaves the place
			break
		}

		// based on current direction next step is going to be an obstacle?
		if in[nextY][nextX] == '#' {
			guardian.Rotate()
			continue
		}

		// move the guardian
		guardian.pos = nextPos
	}

	return false
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	matrix, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(exampleInput))
	if err != nil {
		panic(err)
	}

	pos := FindGuardian(matrix)
	fmt.Printf("Current position of a guardian: %v\n", pos)

	distinctSteps, possibleObstacle := Walk(matrix, pos)
	fmt.Printf("Guarding made %d distinct steps before he left\n", distinctSteps)
	fmt.Printf("We can add %d possible obstacles to make him loop\n", possibleObstacle)
}
