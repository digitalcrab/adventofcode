package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

const (
	Start = 'S'
	End   = 'E'
	Wall  = '#'
)

type State struct {
	pos utils.Pos
	dir int // just an index of Azimuth directions
}

type Step struct {
	state State // this can be integrated but i found it's easier to have this data separated from score
	score int   // used to sort in prio queue
}

func FindBestScorePath(in [][]byte) (int, int) {
	start, end := utils.WhereIs(in, Start), utils.WhereIs(in, End)
	startDirection := 1 // East

	// there is the storage of all positions + direction how we get there
	// and a score
	scores := map[utils.Pos][]int{}
	for p := range utils.Positions(in) {
		// all 4 directions
		scores[p] = make([]int, 4)
		for d := range utils.AzimuthDirections {
			scores[p][d] = math.MaxInt
		}
	}
	// initial start position has a score of 0
	scores[start][startDirection] = 0

	// to reconstruct all shortest paths later, we keep track of predecessors
	predecessors := map[utils.Pos][][]State{}
	for p := range utils.Positions(in) {
		// all 4 directions
		predecessors[p] = make([][]State, 4)
	}

	// start the queue of positions to check
	// push the starting point
	queue := &Queue{}
	heap.Init(queue)
	heap.Push(queue, &Step{
		state: State{
			pos: start,
			dir: startDirection,
		},
		score: 0,
	})

	for queue.Len() > 0 {
		current := heap.Pop(queue).(*Step)
		currentScore := scores[current.state.pos][current.state.dir]

		// check if we are at the end
		if current.state.pos.Eq(end) {
			// for the part 1 we can simply return here
			// return current.score
		}

		// there are 3 possible:
		// 1. move in the given direction by 1 step, with score 1
		{
			nextPos := current.state.pos.Next(utils.AzimuthDirections[current.state.dir])
			if nextY, nextX := nextPos.Values(); in[nextY][nextX] != Wall { // check for a wall
				movementScore := currentScore + 1

				if movementScore < scores[nextPos][current.state.dir] {
					scores[nextPos][current.state.dir] = movementScore

					// found better so we reset the whole thing of predecessors
					predecessors[nextPos][current.state.dir] = []State{current.state}

					heap.Push(queue, &Step{
						state: State{
							pos: nextPos,
							dir: current.state.dir,
						},
						score: movementScore,
					})
				} else if movementScore == scores[nextPos][current.state.dir] {
					// the same score
					predecessors[nextPos][current.state.dir] = append(predecessors[nextPos][current.state.dir], current.state)
				}
			}
		}

		// 2+3. rotate left or right with score 1000
		for _, dx := range []int{1, -1} {
			// works for both -/+ of we add all directions one more time
			newDir := (current.state.dir + dx + len(utils.AzimuthDirections)) % len(utils.AzimuthDirections)
			movementScore := currentScore + 1000
			if movementScore < scores[current.state.pos][newDir] {
				scores[current.state.pos][newDir] = movementScore

				// found better so we reset the whole thing of predecessors
				predecessors[current.state.pos][newDir] = []State{current.state}

				heap.Push(queue, &Step{
					state: State{
						pos: current.state.pos,
						dir: newDir,
					},
					score: movementScore,
				})
			} else if movementScore == scores[current.state.pos][newDir] {
				predecessors[current.state.pos][newDir] = append(predecessors[current.state.pos][newDir], current.state)
			}
		}
	}

	// find the best of all
	bestScore, bestDir := math.MaxInt, -1
	for d, v := range scores[end] {
		if v < bestScore {
			bestScore = v
			bestDir = d
		}
	}

	// end state with the best direction
	endState := State{pos: end, dir: bestDir}

	// todo: state includes direction?!
	visitedStates := make(map[State]bool)
	visitedStates[endState] = true

	// create a stack starting from the end node
	stack := []State{endState}

	// backtrack all nodes and count how many we've visited in all best path's
	for len(stack) > 0 {
		// chop the last
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// for all predecessors
		for _, prevState := range predecessors[current.pos][current.dir] {
			// not seen
			if !visitedStates[prevState] {
				visitedStates[prevState] = true
				// add to stack
				stack = append(stack, prevState)
			}
		}
	}

	// todo: merge with a loop above? maybe?
	visitedNodes := map[utils.Pos]struct{}{}
	for vs := range visitedStates {
		visitedNodes[vs.pos] = struct{}{}
	}

	return bestScore, len(visitedNodes)
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	matrix, _ := utils.ReadFileIntoBytesMatrix(strings.NewReader(exampleInput))
	score, visitedNodes := FindBestScorePath(matrix)
	fmt.Printf("The best score: %d with %d visited nodes\n", score, visitedNodes) // 7036 45
}
