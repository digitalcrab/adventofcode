package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed "input.txt"
var Day5Input string

func TestCalcSummOfMiddlePages(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		args       args
		want       int
		wantSorted int
	}{
		{
			name:       "example",
			args:       args{in: exampleInput},
			want:       143,
			wantSorted: 123,
		},
		{
			name:       "input",
			args:       args{in: Day5Input},
			want:       4790,
			wantSorted: 6319,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules, updates, err := Input(strings.NewReader(tt.args.in))
			if err != nil {
				t.Errorf("unexpected error: %v\n", err)
			}

			correct, incorrect := FilterIncorrectUpdates(rules, updates)
			sorted := SortUpdates(rules, incorrect)

			if got := CalcSummOfMiddlePages(correct); got != tt.want {
				t.Errorf("CalcSummOfMiddlePages(correct) = %v, want %v", got, tt.want)
			}
			if got := CalcSummOfMiddlePages(sorted); got != tt.wantSorted {
				t.Errorf("CalcSummOfMiddlePages(sorted) = %v, want %v", got, tt.wantSorted)
			}
		})
	}
}
