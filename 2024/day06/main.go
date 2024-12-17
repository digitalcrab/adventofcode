package main

import (
	"container/ring"
	_ "embed"
	"fmt"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

type Guardian struct {
	y, x      int
	direction *ring.Ring
}

func (g *Guardian) Direction() *utils.Direction {
	return g.direction.Value.(*utils.Direction)
}

func (g *Guardian) Rotate() {
	g.direction = g.direction.Next()
}

func FindGuardian(in [][]byte) Guardian {
	for y, row := range in {
		for x := range row {
			ch := in[y][x]
			// not a free space and not obstacle
			if ch != '.' && ch != '#' {
				guardian := Guardian{
					y:         y,
					x:         x,
					direction: utils.NewAzimuthRing(utils.SymbolDirection[ch]),
				}
				return guardian
			}
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
	visitedDirections := map[utils.Pos]map[byte]struct{}{}
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
		if pos.Y() == guardian.y && pos.X() == guardian.x {
			continue
		}

		go func() {
			// place a new obstacle
			copyIn := utils.DuplicateBytesMatrix(in)
			copyIn[pos.Y()][pos.X()] = '#'

			// run again the thing
			loopsCh <- walk(copyIn, guardian, map[utils.Pos]map[byte]struct{}{})
		}()
	}

	for pos := range visitedDirections {
		if pos.Y() == guardian.y && pos.X() == guardian.x {
			continue
		}
		if <-loopsCh {
			possibleObstacle++
		}
	}

	return distinctSteps, possibleObstacle
}

func walk(in [][]byte, guardian Guardian, visitedDirections map[utils.Pos]map[byte]struct{}) bool {
	for {
		positionKey := utils.NewPos(guardian.y, guardian.x)

		// record current position + direction
		if _, visited := visitedDirections[positionKey]; !visited {
			visitedDirections[positionKey] = map[byte]struct{}{utils.DirectionSymbol[guardian.Direction()]: {}}
		} else {
			// if visited, we need to detect loop
			dirSym := utils.DirectionSymbol[guardian.Direction()]
			if _, looped := visitedDirections[positionKey][dirSym]; looped {
				return true
			}
			// not looped then just save
			visitedDirections[positionKey][dirSym] = struct{}{}
		}

		// calculate next step as a current step plus directional change
		nextY, nextX := guardian.y+guardian.Direction().Y, guardian.x+guardian.Direction().X

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
		guardian.y, guardian.x = nextY, nextX
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
