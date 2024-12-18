package utils

import "fmt"

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
	for y, row := range in {
		for x := range row {
			if in[y][x] == what {
				return NewPos(y, x)
			}
		}
	}
	return [2]int{}
}
