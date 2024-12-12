package main

import (
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestTotalPrice(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name             string
		args             args
		wantSumPerimeter int
		wantSumSides     int
	}{
		{
			name:             "example",
			args:             args{in: exampleInput},
			wantSumPerimeter: 1930,
			wantSumSides:     1206,
		},
		{
			name:             "input",
			args:             args{in: DayInput},
			wantSumPerimeter: 1304764,
			wantSumSides:     811148,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(tt.args.in)) // 1304712 to low
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			gotSumPerimeter, gotSumSides := TotalPrice(data)
			if gotSumPerimeter != tt.wantSumPerimeter {
				t.Errorf("TotalPrice() gotSumPerimeter = %v, want %v", gotSumPerimeter, tt.wantSumPerimeter)
			}
			if gotSumSides != tt.wantSumSides {
				t.Errorf("TotalPrice() gotSumSides = %v, want %v", gotSumSides, tt.wantSumSides)
			}
		})
	}
}
