package main

import (
	_ "embed"
	"fmt"

	"github.com/digitalcrab/adventofcode/utils"
)

type Robot struct {
	pos      utils.Pos
	velocity utils.Pos
}

func (r *Robot) String() string {
	return fmt.Sprintf("p=%d,%d v=%d,%d", r.pos.X(), r.pos.Y(), r.velocity.X(), r.velocity.Y())
}

func Move(robots []*Robot, steps, height, weight int) {
	for _, r := range robots {
		// as we know number of steps and velocity we can simply
		// multiply steps on velocity and add current coordinates
		// this makes robot far away from the borders (in any of the direction)
		movedY := r.pos.Y() + steps*r.velocity.Y()
		movedX := r.pos.X() + steps*r.velocity.X()

		// if we divide that huge position on the size (height, weight) we get
		// the rest as a new position.
		newY := movedY % height
		newX := movedX % weight

		// thing i've learned hard way:
		// if coordinates are negative then basically % behaves a bit different,
		// it keeps the sign of dividend
		// to make it work, we just add the height or weight one more time and
		// perform % one more time (for positive it does not change anything)
		newY = (newY + height) % height
		newX = (newX + weight) % weight

		r.pos = utils.NewPos(newY, newX)
	}
}

func CountRobotsInQuadrants(robots []*Robot, height, weight int) (q [4]int) {
	midY := (height - 1) / 2
	midX := (weight - 1) / 2
	for _, r := range robots {
		if r.pos.Y() < midY && r.pos.X() < midX { // top left
			q[0]++
		} else if r.pos.Y() < midY && r.pos.X() > midX { // top right
			q[1]++
		} else if r.pos.Y() > midY && r.pos.X() < midX { // bottom left
			q[2]++
		} else if r.pos.Y() > midY && r.pos.X() > midX { // bottom right
			q[3]++
		}
	}
	return
}

func IsThereATree(robots []*Robot) bool {
	// first assumption is that all robots have a unique position
	// so basically all robots in use to display the tree and not a single one is in the same spot
	spots := make(map[utils.Pos]struct{})
	for _, r := range robots {
		if _, seen := spots[r.pos]; seen {
			return false
		}
		spots[r.pos] = struct{}{}
	}
	return true
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	// example
	{
		fmt.Println("Example")
		height := 7
		weight := 11
		robots := Parse(exampleInput)
		Move(robots, 100, height, weight)
		Display(robots, height, weight)
		q := CountRobotsInQuadrants(robots, height, weight)
		factor := q[0] * q[1] * q[2] * q[3]
		fmt.Printf("Safety factor: %d\n", factor) // 12
	}

	// part 1
	{
		fmt.Println("Part 1")
		height := 103
		weight := 101
		robots := Parse(DayInput)
		Move(robots, 100, height, weight)
		Display(robots, height, weight)
		q := CountRobotsInQuadrants(robots, height, weight)
		factor := q[0] * q[1] * q[2] * q[3]
		fmt.Printf("Safety factor: %d\n", factor) // 218433348
	}

	// part 2
	{
		fmt.Println("Part 2")
		height := 103
		weight := 101
		robots := Parse(DayInput)

		step := 1
		for step < 100000 { // just a number with a hope it's going to work
			// move one by one
			Move(robots, 1, height, weight)
			// check if we see a tree
			if IsThereATree(robots) {
				Display(robots, height, weight)
				fmt.Printf("Steps: %d\n", step) // 6512
				break
			}
			// next step
			step++
		}
	}
}
