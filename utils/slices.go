package utils

import (
	"fmt"
	"iter"
)

func DuplicateBytesMatrix(in [][]byte) [][]byte {
	out := make([][]byte, len(in))
	for i := range in {
		out[i] = append(in[i][:0:0], in[i]...)
	}
	return out
}

func PrintMatrix(in [][]byte) {
	for y, row := range in {
		for x := range row {
			fmt.Print(string(in[y][x]))
		}
		fmt.Println()
	}
}

func CreateMatrix(height, width int, ch byte) [][]byte {
	out := make([][]byte, height)
	for y := range out {
		out[y] = make([]byte, width)
		for x := range out[y] {
			out[y][x] = ch
		}
	}
	return out
}

func RepeatInt(value int, count int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = value
	}
	return s
}

func WhereIs(in [][]byte, what byte) Pos {
	for p, ch := range PositionsValues(in) {
		if ch == what {
			return p
		}
	}
	return [2]int{}
}

func Positions(in [][]byte) iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		for y, row := range in {
			for x := range row {
				if !yield(NewPos(y, x)) {
					return
				}
			}
		}
	}
}

func PositionsValues(in [][]byte) iter.Seq2[Pos, byte] {
	return func(yield func(Pos, byte) bool) {
		for y, row := range in {
			for x := range row {
				if !yield(NewPos(y, x), in[y][x]) {
					return
				}
			}
		}
	}
}
