package main

import (
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestSumOfCorrectEquations(t *testing.T) {
	type args struct {
		in        string
		operation []OperationFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{in: exampleInput, operation: []OperationFunc{AdditionOp, MultiplicationOp}},
			want: 3749,
		},
		{
			name: "input",
			args: args{in: DayInput, operation: []OperationFunc{AdditionOp, MultiplicationOp}},
			want: 932137732557,
		},
		{
			name: "example with concat",
			args: args{in: exampleInput, operation: []OperationFunc{AdditionOp, MultiplicationOp, ConcatOp}},
			want: 11387,
		},
		{
			name: "input with concat",
			args: args{in: DayInput, operation: []OperationFunc{AdditionOp, MultiplicationOp, ConcatOp}},
			want: 661823605105500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var equations []Equation
			err := utils.ScanFileLineByLine(strings.NewReader(tt.args.in), func(line string) {
				equations = append(equations, ParseEquation(line))
			})

			if err != nil {
				t.Errorf("unexpected error: %v\n", err)
			}

			if got := SumOfCorrectEquations(equations, tt.args.operation...); got != tt.want {
				t.Errorf("SumOfCorrectEquations() = %v, want %v", got, tt.want)
			}
		})
	}
}
