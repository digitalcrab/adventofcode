package main

import (
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestFindBestScorePath(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
		wantNodes int
	}{
		{
			name:      "example",
			args:      args{in: exampleInput},
			wantScore: 7036,
			wantNodes: 45,
		},
		{
			name:      "input",
			args:      args{in: DayInput},
			wantScore: 103512,
			wantNodes: 554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix, _ := utils.ReadFileIntoBytesMatrix(strings.NewReader(tt.args.in))
			gotScore, gotNodes := FindBestScorePath(matrix)
			if gotScore != tt.wantScore {
				t.Errorf("FindBestScorePath() = %v, wantScore %v", gotScore, tt.wantScore)
			}
			if gotNodes != tt.wantNodes {
				t.Errorf("FindBestScorePath() = %v, wantNodes %v", gotNodes, tt.wantNodes)
			}
		})
	}
}
