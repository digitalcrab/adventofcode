package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		args   args
		want   int
		wantV2 int
	}{
		{
			name:   "example",
			args:   args{in: exampleInput},
			want:   1928,
			wantV2: 2858,
		},
		{
			name:   "input",
			args:   args{in: DayInput},
			want:   6307275788409,
			wantV2: 6327174563252,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialBlocks, freeSpaceHeaps := ReadDiskBlocks(tt.args.in)
			movedBlockes := MoveFreeSpace(initialBlocks)
			movedBlockes2 := MoveFreeSpaceV2(initialBlocks, freeSpaceHeaps)
			if got := Checksum(movedBlockes); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
			if got := Checksum(movedBlockes2); got != tt.wantV2 {
				t.Errorf("ChecksumV2() = %v, want %v", got, tt.wantV2)
			}
		})
	}
}
