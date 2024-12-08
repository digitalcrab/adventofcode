package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

// Equation represent set of operands and possible result
// Operators are always evaluated left-to-right, not according to precedence rules.
// Furthermore, numbers in the equations cannot be rearranged.
// add (+) and multiply (*)
type Equation struct {
	Result   int
	Operands []int
}

type OperationFunc func(a, b int) int

func AdditionOp(a, b int) int {
	return a + b
}

func MultiplicationOp(a, b int) int {
	return a * b
}

func ConcatOp(a, b int) int {
	return utils.Int(fmt.Sprintf("%d%d", a, b))
}

func (e Equation) Evaluate(operations ...OperationFunc) bool {
	// basically invalid
	if len(e.Operands) == 0 {
		return false
	}

	// queue of prev results, start with a first operand
	queue := []int{e.Operands[0]}

	// loop through the rest of operands
	for _, number := range e.Operands[1:] {
		// new results we store here
		var newResults []int

		// for every prev result apply every operation
		for _, prevResult := range queue {
			for _, op := range operations {
				// calculate the result
				intermediateRes := op(prevResult, number)
				// store in the new results
				newResults = append(newResults, intermediateRes)
			}
		}

		// replace the queue with new results
		queue = newResults
	}

	// now we have all the results in the queue
	// check if we have the one we need

	for _, res := range queue {
		if res == e.Result {
			return true
		}
	}

	return false
}

// ParseEquation parses the string into Equation
func ParseEquation(s string) Equation {
	parts := strings.SplitN(s, ":", 2)
	e := Equation{
		Result: utils.Int(parts[0]),
	}
	for _, sn := range strings.Split(strings.TrimSpace(parts[1]), " ") {
		e.Operands = append(e.Operands, utils.Int(sn))
	}
	return e
}

func SumOfCorrectEquations(equations []Equation, operations ...OperationFunc) int {
	var sum int
	for _, e := range equations {
		if e.Evaluate(operations...) {
			sum += e.Result
		}
	}
	return sum
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	var equations []Equation
	err := utils.ScanFileLineByLine(strings.NewReader(exampleInput), func(line string) {
		equations = append(equations, ParseEquation(line))
	})

	if err != nil {
		panic(err)
	}

	sum1 := SumOfCorrectEquations(equations, AdditionOp, MultiplicationOp)
	fmt.Printf("The sum of all correct equasions is %d (based on + and *)\n", sum1)

	sum2 := SumOfCorrectEquations(equations, AdditionOp, MultiplicationOp, ConcatOp)
	fmt.Printf("The sum of all correct equasions is %d (based on + and * and ||)\n", sum2)
}
