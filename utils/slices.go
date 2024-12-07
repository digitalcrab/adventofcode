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
	for rx, row := range in {
		for cx := range row {
			fmt.Print(string(in[rx][cx]))
		}
		fmt.Println()
	}
}
