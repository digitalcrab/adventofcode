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
	predecessors := map[utils.Pos][][]utils.StepState{}
	for p := range utils.Positions(in) {
		// all 4 directions
		predecessors[p] = make([][]utils.StepState, 4)
	}

	// start the queue of positions to check
	// push the starting point
	queue := &utils.Queue{}
	heap.Init(queue)
	heap.Push(queue, &utils.Step{
		St: utils.StepState{
			P: start,
			D: startDirection,
		},
		S: 0,
	})

	for queue.Len() > 0 {
		current := heap.Pop(queue).(*utils.Step)
		currentScore := scores[current.St.P][current.St.D]

		// check if we are at the end
		if current.St.P.Eq(end) {
			// for the part 1 we can simply return here
			// return current.score
		}

		// there are 3 possible:
		// 1. move in the given direction by 1 step, with score 1
		{
			nextPos := current.St.P.Next(utils.AzimuthDirections[current.St.D])
			if nextY, nextX := nextPos.Values(); in[nextY][nextX] != Wall { // check for a wall
				movementScore := currentScore + 1

				if movementScore < scores[nextPos][current.St.D] {
					scores[nextPos][current.St.D] = movementScore

					// found better so we reset the whole thing of predecessors
					predecessors[nextPos][current.St.D] = []utils.StepState{current.St}

					heap.Push(queue, &utils.Step{
						St: utils.StepState{
							P: nextPos,
							D: current.St.D,
						},
						S: movementScore,
					})
				} else if movementScore == scores[nextPos][current.St.D] {
					// the same score
					predecessors[nextPos][current.St.D] = append(predecessors[nextPos][current.St.D], current.St)
				}
			}
		}

		// 2+3. rotate left or right with score 1000
		for _, dx := range []int{1, -1} {
			// works for both -/+ of we add all directions one more time
			newDir := (current.St.D + dx + len(utils.AzimuthDirections)) % len(utils.AzimuthDirections)
			movementScore := currentScore + 1000
			if movementScore < scores[current.St.P][newDir] {
				scores[current.St.P][newDir] = movementScore

				// found better so we reset the whole thing of predecessors
				predecessors[current.St.P][newDir] = []utils.StepState{current.St}

				heap.Push(queue, &utils.Step{
					St: utils.StepState{
						P: current.St.P,
						D: newDir,
					},
					S: movementScore,
				})
			} else if movementScore == scores[current.St.P][newDir] {
				predecessors[current.St.P][newDir] = append(predecessors[current.St.P][newDir], current.St)
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
	endState := utils.StepState{P: end, D: bestDir}

	// todo: state includes direction?!
	visitedStates := make(map[utils.StepState]bool)
	visitedStates[endState] = true

	// create a stack starting from the end node
	stack := []utils.StepState{endState}

	// backtrack all nodes and count how many we've visited in all best path's
	for len(stack) > 0 {
		// chop the last
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// for all predecessors
		for _, prevState := range predecessors[current.P][current.D] {
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
		visitedNodes[vs.P] = struct{}{}
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
