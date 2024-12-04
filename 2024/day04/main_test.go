package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/digitalcrab/adventofcode/utils"
)

//go:embed "input.txt"
var Day4Input string

func TestCountXMAS(t *testing.T) {
	type args struct {
		in   string
		word []byte
	}
	tests := []struct {
		name       string
		args       args
		wantXMAS   int
		wantXXXMAS int
	}{
		{
			name:       "example",
			args:       args{in: exampleInput},
			wantXMAS:   18,
			wantXXXMAS: 9,
		},
		{
			name:       "input",
			args:       args{in: Day4Input},
			wantXMAS:   2507,
			wantXXXMAS: 1969, // 2025 if MAM and SAS counts, maaan ;)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := utils.ReadFileIntoBytesMatrix(strings.NewReader(tt.args.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got := CountXMAS(data); got != tt.wantXMAS {
				t.Errorf("CountXMAS() = %v, want %v", got, tt.wantXMAS)
			}

			if got := CountXXXMAS(data); got != tt.wantXXXMAS {
				t.Errorf("CountXXXMAS() = %v, want %v", got, tt.wantXXXMAS)
			}
		})
	}
}
