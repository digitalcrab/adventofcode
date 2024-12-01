package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDistanceBetweenLists(t *testing.T) {
	task1in1, task1in2 := input1(t)

	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name           string
		args           args
		wantDistance   int
		wantSimilarity int
	}{
		{
			name: "example",
			args: args{
				l1: []int{3, 4, 2, 1, 3, 3},
				l2: []int{4, 3, 5, 3, 9, 3},
			},
			wantDistance:   11,
			wantSimilarity: 31,
		},
		{
			name: "first task",
			args: args{
				l1: task1in1,
				l2: task1in2,
			},
			wantDistance:   2066446,
			wantSimilarity: 24931009,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistanceBetweenLists(tt.args.l1, tt.args.l2); got != tt.wantDistance {
				t.Errorf("DistanceBetweenLists() = %v, want %v", got, tt.wantDistance)
			}
			if got := SimilarityScore(tt.args.l1, tt.args.l2); got != tt.wantSimilarity {
				t.Errorf("SimilarityScore() = %v, want %v", got, tt.wantSimilarity)
			}
		})
	}
}

func input1(t *testing.T) ([]int, []int) {
	t.Helper()

	readFile, err := os.Open("input1.txt")
	if err != nil {
		t.Fatalf("can not open input file: %v", err)
	}
	defer func() {
		_ = readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var l1, l2 []int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		numbers := strings.SplitN(line, "   ", 2)
		if len(numbers) != 2 {
			t.Fatalf("unexpected count of numbers %d in the line %s", len(numbers), line)
		}

		l1 = append(l1, number(t, numbers[0]))
		l2 = append(l2, number(t, numbers[1]))
	}

	return l1, l2
}

func number(t *testing.T, s string) int {
	t.Helper()

	num, err := strconv.Atoi(s)
	if err != nil {
		t.Fatalf("can not parse number from the string %q: %v", s, err)
	}

	return num
}
