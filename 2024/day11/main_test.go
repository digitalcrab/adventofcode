package main

import "testing"

func TestBlinkTimes(t *testing.T) {
	type args struct {
		lineOfStones string
		times        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{lineOfStones: exampleInput, times: 25},
			want: 55312,
		},
		{
			name: "input",
			args: args{lineOfStones: DayInput, times: 25},
			want: 183484,
		},
		{
			name: "input",
			args: args{lineOfStones: DayInput, times: 75},
			want: 218817038947400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlinkTimes(tt.args.lineOfStones, tt.args.times); got != tt.want {
				t.Errorf("BlinkTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
