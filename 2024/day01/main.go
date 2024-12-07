package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"

	"github.com/digitalcrab/adventofcode/utils"
)

// DistanceBetweenLists calculates the distance between lists.
// To find out, pair up the numbers and measure how far apart they are.
// Pair up the smallest number in the left list with the smallest number in the right list,
// then the second-smallest left number with the second-smallest right number, and so on.
//
// Within each pair, figure out how far apart the two numbers are;
// you'll need to add up all of those distances. For example, if you pair up
// a 3 from the left list with a 7 from the right list, the distance apart is 4;
// if you pair up a 9 with a 3, the distance apart is 6.
func DistanceBetweenLists(l1, l2 []int) int {
	// The best approach is to use min-heap for each list.
	// Populate them and calculate distance.

	l1h := utils.InitIntHeap(l1)
	l2h := utils.InitIntHeap(l2)

	var distance int

	// Loop over l2
	for l2h.Len() > 0 {
		el1 := heap.Pop(&l1h).(int)
		el2 := heap.Pop(&l2h).(int)

		// Distance could be negative, but we actually need only number without the sign
		d := math.Abs(float64(el2 - el1))

		distance += int(d)
	}

	return distance
}

// SimilarityScore calculates the score.
// This time, you'll need to figure out exactly how often each number from the left list
// appears in the right list. Calculate a total similarity score by adding up each number
// in the left list after multiplying it by the number of times that number appears in
// the right list.
func SimilarityScore(l1, l2 []int) int {
	// First lets find how many times each number appear in the list.
	times1 := calculateTimes(l1)
	times2 := calculateTimes(l2)

	var similarity int

	for num, leftTimes := range times1 {
		rightTimes := times2[num]
		similarity += num * rightTimes * leftTimes
	}

	return similarity
}

func calculateTimes(l []int) map[int]int {
	m := make(map[int]int)
	for _, n := range l {
		if times, exist := m[n]; exist {
			m[n] = times + 1
		} else {
			m[n] = 1
		}
	}
	return m
}

//go:embed "input.txt"
var DayInput string

func main() {
	l1 := []int{3, 4, 2, 1, 3, 3}
	l2 := []int{4, 3, 5, 3, 9, 3}
	distance := DistanceBetweenLists(l1, l2)
	fmt.Println(distance) // Output: 11
	similarity := SimilarityScore(l1, l2)
	fmt.Println(similarity) // Output: 31
}
