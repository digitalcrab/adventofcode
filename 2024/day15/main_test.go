package main

import (
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestMoveGPS(t *testing.T) {
	type args struct {
		inMap, inMoves string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "example",
			args:    args{inMap: exampleInputMap, inMoves: exampleInputMove},
			wantSum: 10092,
		},
		{
			name:    "input",
			args:    args{inMap: DayInputMap, inMoves: DayInputMove},
			wantSum: 1514333,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, moves := Map(tt.args.inMap), Movements(tt.args.inMoves)
			robot := WhereIsRobot(matrix)

			for _, ch := range moves {
				matrix, robot = Move(matrix, robot, utils.SymbolDirection[ch])
			}

			if gotSum := GPS(matrix, Box); gotSum != tt.wantSum {
				t.Errorf("GPS() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestMove2GPS(t *testing.T) {
	type args struct {
		inMap, inMoves string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "example",
			args:    args{inMap: exampleInputMap, inMoves: exampleInputMove},
			wantSum: 9021,
		},
		{
			name:    "input",
			args:    args{inMap: DayInputMap, inMoves: DayInputMove},
			wantSum: 1528453,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, moves := Map(tt.args.inMap), Movements(tt.args.inMoves)
			matrix = Enlarge(matrix)
			robot := WhereIsRobot(matrix)

			for _, ch := range moves {
				matrix, robot = Move2(matrix, robot, utils.SymbolDirection[ch])
			}

			if gotSum := GPS(matrix, BoxLeft); gotSum != tt.wantSum {
				t.Errorf("GPS() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
