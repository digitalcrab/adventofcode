package main

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/digitalcrab/adventofcode/utils"
)

const (
	Wall     = '#'
	Box      = 'O'
	BoxLeft  = '['
	BoxRight = ']'
	Robot    = '@'
	Space    = '.'
)

func WhereIsRobot(in [][]byte) utils.Pos {
	for y, row := range in {
		for x := range row {
			if in[y][x] == Robot {
				return utils.NewPos(y, x)
			}
		}
	}
	return [2]int{}
}

func Move(in [][]byte, robot utils.Pos, direction *utils.Direction) ([][]byte, utils.Pos) {
	nextY := robot.Y() + direction.Y
	nextX := robot.X() + direction.X

main:
	switch nextCh := in[nextY][nextX]; nextCh {
	case Wall: // next step is a wall, do nothing
		break
	case Space:
		// switch the space with robot
		in[robot.Y()][robot.X()], in[nextY][nextX] = Space, Robot
		// move robot
		robot = utils.NewPos(nextY, nextX)
	case Box: // part 1
		// if we see the box, we need to remember its position
		boxes := []utils.Pos{
			utils.NewPos(nextY, nextX),
		}
		boxesY, boxesX := nextY, nextX
		// and try to move with the box
		for {
			boxesY = boxesY + direction.Y
			boxesX = boxesX + direction.X
			// hit the wall somewhere along the way? no movements possible
			if in[boxesY][boxesX] == Wall {
				break main
			}
			// another box? remember it, so we can push it later
			if in[boxesY][boxesX] == Box {
				boxes = append(boxes, utils.NewPos(boxesY, boxesX))
				continue
			}
			// found the first free space, need to push all boxes here
			if in[boxesY][boxesX] == Space {
				// move all boxes (starting from the back)
				for _, b := range slices.Backward(boxes) {
					in[b.Y()][b.X()], in[b.Y()+direction.Y][b.X()+direction.X] = Space, Box
				}
				// switch the space with robot
				in[robot.Y()][robot.X()], in[nextY][nextX] = Space, Robot
				// move robot
				robot = utils.NewPos(nextY, nextX)
				break main
			}
		}
	}

	return in, robot
}

func Move2(in [][]byte, robot utils.Pos, direction *utils.Direction) ([][]byte, utils.Pos) {
	// here i am trying to rethink what `Move` as doing,
	// instead of working on the `ch` bases, i'll try to move more stuff at once

	// start with list of things to move, initially only robot
	queue := []utils.Pos{
		robot,
	}
	idx := 0

	// this is going to be just a small safety measure not to
	// add the same and the same thing again in the queue
	seen := make(map[utils.Pos]struct{}) // todo: a bit ugly

	// we are checking item by item
	for idx < len(queue) {
		// get the next item
		current := queue[idx]
		// move it
		nextY := current.Y() + direction.Y
		nextX := current.X() + direction.X
		nextPos := utils.NewPos(nextY, nextX)

		ch := in[nextY][nextX]

		// if we hit the wall anywhere on the way (by the robot or a box)
		// no movements possible
		if ch == Wall {
			return in, robot
		}

		if ch == BoxLeft || ch == BoxRight {
			// store basically the point to move, and anchor the next step to check
			if _, ok := seen[nextPos]; !ok {
				queue = append(queue, nextPos)
				seen[nextPos] = struct{}{}
			}
			// if we see the left part of the box, then add right part as well
			if ch == BoxLeft {
				rightPart := utils.NewPos(nextY, nextX+1)
				if _, ok := seen[rightPart]; !ok {
					queue = append(queue, rightPart)
					seen[rightPart] = struct{}{}
				}
			} else {
				leftPart := utils.NewPos(nextY, nextX-1)
				if _, ok := seen[leftPart]; !ok {
					queue = append(queue, leftPart)
					seen[leftPart] = struct{}{}
				}
			}
		}

		// move to the next item in the queue
		idx++
	}

	// here we should have all the items that needs to move collected in `queue`
	// not sure about the order of that stuff, hmmm ... ?
	for _, prev := range slices.Backward(queue) {
		// replace next position with prev position value
		// prev replace with space
		in[prev.Y()+direction.Y][prev.X()+direction.X], in[prev.Y()][prev.X()] = in[prev.Y()][prev.X()], Space
	}

	// move robot
	robot = utils.NewPos(robot.Y()+direction.Y, robot.X()+direction.X)

	return in, robot
}

func GPS(in [][]byte, ch byte) (sum int) {
	for y, row := range in {
		for x := range row {
			if in[y][x] == ch {
				sum += 100*y + x
			}
		}
	}
	return
}

func Enlarge(in [][]byte) [][]byte {
	bigger := make([][]byte, len(in)) // rows are the same
	for y, row := range in {
		bigger[y] = make([]byte, 0, len(row)*2) // 2 times wider
		for x := range row {
			if in[y][x] == Wall {
				bigger[y] = append(bigger[y], Wall, Wall)
			}
			if in[y][x] == Box {
				bigger[y] = append(bigger[y], BoxLeft, BoxRight)
			}
			if in[y][x] == Space {
				bigger[y] = append(bigger[y], Space, Space)
			}
			if in[y][x] == Robot {
				bigger[y] = append(bigger[y], Robot, Space)
			}
		}
	}
	return bigger
}

//go:embed "example_map.txt"
var exampleInputMap string

//go:embed "example_move.txt"
var exampleInputMove string

//go:embed "input_map.txt"
var DayInputMap string

//go:embed "input_move.txt"
var DayInputMove string

func main() {
	matrix, moves := Map(exampleInputMap), Movements(exampleInputMove)
	robot := WhereIsRobot(matrix)
	fmt.Printf("Robot location: %v\n", robot)
	utils.PrintMatrix(matrix)

	for _, ch := range moves {
		fmt.Printf("Move %s:\n", string(ch))
		matrix, robot = Move(matrix, robot, utils.SymbolDirection[ch])
		utils.PrintMatrix(matrix)
	}

	// Calculate GPS
	sum := GPS(matrix, Box)
	fmt.Printf("Sum of all boxes' GPS coordinates: %d\n", sum) // 10092

	fmt.Println("..... Enlarge .....")
	matrix = Map(exampleInputMap)
	matrix = Enlarge(matrix)
	robot = WhereIsRobot(matrix)
	fmt.Printf("Robot location: %v\n", robot)
	utils.PrintMatrix(matrix)

	for _, ch := range moves {
		fmt.Printf("Move %s:\n", string(ch))
		matrix, robot = Move2(matrix, robot, utils.SymbolDirection[ch])
		utils.PrintMatrix(matrix)
	}

	// Calculate GPS
	sum = GPS(matrix, BoxLeft)
	fmt.Printf("Sum of all boxes' GPS coordinates: %d\n", sum) // 9021

	// 1528453
}
