package main

import "testing"

func TestAllPricesQueue(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "example",
			args:    args{in: exampleInput},
			wantSum: 480,
		},
		{
			name:    "input",
			args:    args{in: DayInput},
			wantSum: 29522, // this is already slow
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machines := Parse(tt.args.in)
			if gotSum := AllPricesQueue(machines); gotSum != tt.wantSum {
				t.Errorf("AllPricesQueue() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestAllPricesMath(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "input",
			args:    args{in: DayInput},
			wantSum: 101214869433312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machines := Parse(tt.args.in)
			if gotSum := AllPricesMath(machines, 10000000000000); gotSum != tt.wantSum {
				t.Errorf("AllPricesQueue() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
