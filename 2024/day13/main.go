package main

import (
	"container/heap"
	_ "embed"
	"fmt"

	"github.com/digitalcrab/adventofcode/utils"
)

type Button struct {
	move  utils.Pos
	price int
}

type Machine struct {
	buttons []Button
	price   utils.Pos
}

// CostMath solves the part 2 problem, and part 1 as well with simple math ;)
// This is basically equation with 2 unknowns:
//
//	AT*A.x + BT*B.x = P.x
//	AT*A.y + BT*B.y = P.y
//
// Where:
//
//	P.x and P.y - destination of a prize
//	A.x and A.y - how far A button moves us closer
//	B.x and B.y - same for button B
//	AT and BT - how many times we push the button (A and B respectfully)
//
// This makes:
//
//	 AT = (P.x*B.y - P.y*B.x) / (B.y*A.x - A.y*B.x)
//		BT = (P.y*A.x - P.x*A.y) / (B.y*A.x - A.y*B.x)
//
// See my funny image with old-school calculations on the paper ;) -  day13.jpg
func (m Machine) CostMath(multiplier int) int {
	// both buttons
	buttonA := m.buttons[0]
	buttonB := m.buttons[1]

	// multiply the price
	prizeY := m.price.Y() + multiplier
	prizeX := m.price.X() + multiplier

	// calculate the common divider
	divider := buttonB.move.Y()*buttonA.move.X() - buttonA.move.Y()*buttonB.move.X()

	at := (prizeX*buttonB.move.Y() - prizeY*buttonB.move.X()) / divider
	bt := (prizeY*buttonA.move.X() - prizeX*buttonA.move.Y()) / divider

	// check the equation
	if (at*buttonA.move.X()+bt*buttonB.move.X() == prizeX) &&
		(at*buttonA.move.Y()+bt*buttonB.move.Y() == prizeY) {
		return at*buttonA.price + bt*buttonB.price
	}

	// not valid
	return -1
}

// Cost solves the part 1 problem by creating a min-heap and trying to each the best price
func (m Machine) Cost() int {
	// create a simple min heap that is going to have the cheapest option easily available
	var queue Steps
	heap.Init(&queue)

	// push the starting point
	heap.Push(&queue, Step{cost: 0, y: 0, x: 0})

	// storage for the point (Y,X) with the price we've used to achieve it
	pricePoint := map[utils.Pos]int{
		utils.NewPos(0, 0): 0,
	}

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(Step)

		// check if we reached the target, return to total cost
		if current.y == m.price[0] && current.x == m.price[1] {
			return current.cost
		}
		// if we way over the target, nothing left to do
		if current.y > m.price[0] || current.x > m.price[1] {
			continue
		}

		// press button A and B
		for _, btn := range m.buttons {
			pressButton(current, btn, pricePoint, &queue)
		}
	}

	// No solution found
	return -1
}

func pressButton(current Step, btn Button, pricePoint map[utils.Pos]int, queue *Steps) {
	nextY := current.y + btn.move.Y()
	nextX := current.x + btn.move.X()
	nextPos := utils.NewPos(nextY, nextX)
	nextCost := current.cost + btn.price

	// If we haven't visited this position or found a cheaper cost, update and push
	if knownPrice, ok := pricePoint[nextPos]; !ok || nextCost < knownPrice {
		pricePoint[nextPos] = nextCost
		heap.Push(queue, Step{cost: nextCost, y: nextY, x: nextX})
	}
}

func AllPricesQueue(machines []Machine) (sum int) {
	for _, m := range machines {
		cost := m.Cost()
		if cost == -1 {
			continue
		}
		sum += cost
	}
	return
}

func AllPricesMath(machines []Machine, multiplier int) (sum int) {
	for _, m := range machines {
		cost := m.CostMath(multiplier)
		if cost == -1 {
			continue
		}
		sum += cost
	}
	return
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	machines := Parse(DayInput)
	sum := AllPricesMath(machines, 10000000000000)
	fmt.Println(sum)
}
