package main

import (
	_ "embed"
	"testing"
)

//go:embed "input.txt"
var Day2Input string

func TestNumberOfSafeReports(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name          string
		args          args
		want          int
		wantTolerance int
	}{
		{
			name:          "example",
			args:          args{input: exampleInput},
			want:          2,
			wantTolerance: 4,
		},
		{
			name:          "input",
			args:          args{input: Day2Input},
			want:          472,
			wantTolerance: 520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotTolerance := NumberOfSafeReports(tt.args.input, false)
			if got != tt.want {
				t.Errorf("NumberOfSafeReports() got = %v, want %v", got, tt.want)
			}
			if gotTolerance != tt.wantTolerance {
				t.Errorf("NumberOfSafeReports() gotTolerance = %v, wantTolerance %v", gotTolerance, tt.wantTolerance)
			}
		})
	}
}
