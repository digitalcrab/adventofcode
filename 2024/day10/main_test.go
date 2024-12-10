package main

import (
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestSumOfTrailheadScores(t *testing.T) {
	type args struct {
		in         string
		uniqueWays bool
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "example unique",
			args:    args{in: exampleInput, uniqueWays: true},
			wantSum: 36,
		},
		{
			name:    "example",
			args:    args{in: exampleInput, uniqueWays: false},
			wantSum: 81,
		},
		{
			name:    "input unique",
			args:    args{in: DayInput, uniqueWays: true},
			wantSum: 617,
		},
		{
			name:    "input",
			args:    args{in: DayInput, uniqueWays: false},
			wantSum: 1477,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(tt.args.in))
			if err != nil {
				panic(err)
			}

			if gotSum := SumOfTrailheadScores(data, tt.args.uniqueWays); gotSum != tt.wantSum {
				t.Errorf("SumOfTrailheadScores() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
