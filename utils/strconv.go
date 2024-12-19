package utils

import "strconv"

func StringsToInts(ss []string) []int {
	ints := make([]int, len(ss))
	for i, s := range ss {
		ints[i] = Int(s)
	}
	return ints
}

func IntsToStrings(ints []int) []string {
	ss := make([]string, len(ints))
	for i, s := range ints {
		ss[i] = strconv.Itoa(s)
	}
	return ss
}

func Int(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
