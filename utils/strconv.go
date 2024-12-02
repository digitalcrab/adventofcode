package utils

import "strconv"

func Integers(ss []string) []int {
	ints := make([]int, len(ss))
	for i, s := range ss {
		ints[i] = Int(s)
	}
	return ints
}

func Int(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
