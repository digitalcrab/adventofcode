package main

import (
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestCalculateUniqueAntiNodes(t *testing.T) {
	type args struct {
		in       string
		distance float64
		limit    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{in: exampleInput, distance: 2.0, limit: 1},
			want: 14,
		},
		{
			name: "example unlimited same distance",
			args: args{in: exampleInput, distance: 1.0, limit: -1},
			want: 34,
		},
		{
			name: "input",
			args: args{in: DayInput, distance: 2.0, limit: 1},
			want: 269,
		},
		{
			name: "input unlimited same distance",
			args: args{in: DayInput, distance: 1.0, limit: -1},
			want: 949,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(tt.args.in))
			if err != nil {
				panic(err)
			}
			if got := CalculateUniqueAntiNodes(matrix, tt.args.distance, tt.args.limit); got != tt.want {
				t.Errorf("CalculateUniqueAntiNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
