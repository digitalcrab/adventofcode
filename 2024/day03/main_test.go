package main

import (
	_ "embed"
	"testing"
)

//go:embed "input.txt"
var Day3Input string

func TestDecodeMemory(t *testing.T) {
	type args struct {
		data    string
		doLogic bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example (doLogic = false)",
			args: args{data: exampleInput, doLogic: false},
			want: 161,
		},
		{
			name: "example (doLogic = true)",
			args: args{data: exampleInput, doLogic: true},
			want: 161,
		},
		{
			name: "input (doLogic = false)",
			args: args{data: Day3Input, doLogic: false},
			want: 179834255,
		},
		{
			name: "input (doLogic = true)",
			args: args{data: Day3Input, doLogic: true},
			want: 80570939,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeMemory(tt.args.data, tt.args.doLogic); got != tt.want {
				t.Errorf("DecodeMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
