package main

import (
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

func TestMachine_Run(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name:    "example",
			args:    args{in: exampleInput},
			wantOut: "5,7,3,0",
		},
		{
			name:    "input",
			args:    args{in: DayInput},
			wantOut: "5,0,3,5,7,6,1,5,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, p := Parse(tt.args.in)
			out := m.Run(p)
			gotOut := strings.Join(utils.IntsToStrings(out), ",")
			if gotOut != tt.wantOut {
				t.Errorf("Run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestFindBackwards(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "input",
			args: args{in: DayInput},
			want: 164516454365621,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, p := Parse(tt.args.in)
			if got := FindBackwards(p, 0); got != tt.want {
				t.Errorf("FindBackwards() = %v, want %v", got, tt.want)
			}
		})
	}
}
