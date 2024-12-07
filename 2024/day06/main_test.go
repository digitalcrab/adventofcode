package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestWalk(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name                 string
		args                 args
		wantDistinctSteps    int
		wantPossibleObstacle int
	}{
		{
			name:                 "example",
			args:                 args{in: exampleInput},
			wantDistinctSteps:    41,
			wantPossibleObstacle: 6,
		},
		{
			name:                 "input",
			args:                 args{in: DayInput},
			wantDistinctSteps:    5162,
			wantPossibleObstacle: 1909, //
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(tt.args.in))
			if err != nil {
				panic(err)
			}

			pos := FindGuardian(matrix)
			gotDistinctSteps, gotPossibleObstacle := Walk(matrix, pos)
			if gotDistinctSteps != tt.wantDistinctSteps {
				t.Errorf("Walk() DistinctSteps = %v, want %v", gotDistinctSteps, tt.wantDistinctSteps)
			}
			if gotPossibleObstacle != tt.wantPossibleObstacle {
				t.Errorf("Walk() PossibleObstacle = %v, want %v", gotPossibleObstacle, tt.wantPossibleObstacle)
			}
		})
	}
}
