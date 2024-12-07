package main

import (
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

// OrderRule is a map of page and all the pages that need to be after or before this page
type OrderRule struct {
	After  map[string]map[string]struct{}
	Before map[string]map[string]struct{}
}

// Update is a list of page number
type Update []string

func IsInCorrectOrder(rules OrderRule, update Update) bool {
	var seenPages []string

	for _, page := range update {
		// get the rule that's says which pages has to be before
		needBefore := rules.Before[page]

		// if we need something to be before but based on the rules the is nothing defined
		// then for sure incorrect
		if len(needBefore) == 0 && len(seenPages) > 0 {
			return false
		}

		// loop over all seen pages and verify that they are indeed needs to be before
		for _, seen := range seenPages {
			if _, exists := needBefore[seen]; !exists {
				return false
			}
		}

		// add page to the seen slice
		seenPages = append(seenPages, page)
	}

	return true
}

func FilterIncorrectUpdates(rules OrderRule, updates []Update) ([]Update, []Update) {
	incorrect := make([]Update, 0, len(updates))
	correct := make([]Update, 0, len(updates))
	for _, update := range updates {
		if IsInCorrectOrder(rules, update) {
			correct = append(correct, update)
		} else {
			incorrect = append(incorrect, update)
		}
	}
	return correct, incorrect
}

func CalcSummOfMiddlePages(updates []Update) int {
	var sum int
	for _, update := range updates {
		midIdx := (len(update) - 1) / 2
		sum += utils.Int(update[midIdx])
	}
	return sum
}

func SortUpdate(rules OrderRule, update Update) Update {
	sort.Slice(update, func(i, j int) bool {
		needAfter, afterExists := rules.After[update[i]]
		// if there is nothing after that item it's bigger
		if !afterExists {
			return false
		}
		// if j needs to be after i
		_, exists := needAfter[update[j]]
		return exists
	})
	return update
}

func SortUpdates(rules OrderRule, updates []Update) []Update {
	sorted := make([]Update, len(updates))
	for i, update := range updates {
		sorted[i] = SortUpdate(rules, update)
	}
	return sorted
}

func Input(file io.Reader) (OrderRule, []Update, error) {
	rules := OrderRule{
		After:  make(map[string]map[string]struct{}),
		Before: make(map[string]map[string]struct{}),
	}
	updates := make([]Update, 0)
	return rules, updates, utils.ScanFileLineByLine(file, func(line string) {
		if line == "" {
			return
		}

		// this is the order rule
		if strings.Contains(line, "|") {
			orderPages := strings.SplitN(line, "|", 2)
			before, after := orderPages[0], orderPages[1]

			if _, exists := rules.After[before]; !exists {
				rules.After[before] = make(map[string]struct{})
			}
			rules.After[before][after] = struct{}{}

			if _, exists := rules.Before[after]; !exists {
				rules.Before[after] = make(map[string]struct{})
			}
			rules.Before[after][before] = struct{}{}
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	})
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	rules, updates, err := Input(strings.NewReader(exampleInput))
	if err != nil {
		panic(err)
	}

	correct, incorrect := FilterIncorrectUpdates(rules, updates)
	fmt.Printf("Correct updates: %v\n", correct)
	fmt.Printf("Incorrect updates: %v\n", incorrect)

	sorted := SortUpdates(rules, incorrect)
	fmt.Printf("Sorted updates: %v\n", sorted)

	sum := CalcSummOfMiddlePages(correct)
	fmt.Printf("Summ of correct updates middle elements: %d\n", sum)

	sortedSum := CalcSummOfMiddlePages(sorted)
	fmt.Printf("Summ of sorted updates middle elements: %d\n", sortedSum)
}
