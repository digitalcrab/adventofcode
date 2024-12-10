package main

import (
	"bytes"
	"container/heap"
	_ "embed"
	"fmt"
	"math"

	"github.com/digitalcrab/adventofcode/utils"
)

// DiskBlocks represent a slice of dist elements, where -1 means free space
type DiskBlocks []int

func (numbers DiskBlocks) String() string {
	buf := &bytes.Buffer{}
	for _, num := range numbers {
		char := "."
		if num >= 0 {
			char = fmt.Sprintf("%d", num)
		}
		buf.WriteString(char)
	}
	return buf.String()
}

func ReadDiskBlocks(line string) (expanded DiskBlocks, freeSpaceHeaps [10]utils.IntHeap) {
	// here we're going to store the starting points of a free spaces split by its size
	// so basically all free space of size 3 end up in freeSpaceHeaps[3], and as it's heap
	// it's going to order from smallest (earliest point) to the oldest

	idx := 0

	for i, ch := range line {
		// size of the block, all possible values are 0-9
		size := int(ch - '0')

		if i%2 == 0 { // even, means file
			expanded = append(expanded, utils.RepeatInt(idx, size)...)
			// increase index of a file when we've done processing
			idx++
		} else { // odd, free space, we use -1 as a value
			// todo: size == 0 ?? does nothing but maybe to skip it

			// we add to the heap of `size` new point, that has to be the starting point
			// of that block, in our case this is the size of `expanded` before we add -1 to it
			freeSpaceHeaps[size] = append(freeSpaceHeaps[size], len(expanded))

			// add items of a free space
			expanded = append(expanded, utils.RepeatInt(-1, size)...)
		}
	}

	// init all heaps once, it's basically important and we get heap sorted
	for i := range freeSpaceHeaps {
		freeSpaceHeaps[i] = utils.InitIntHeap(freeSpaceHeaps[i])
	}

	return
}

func MoveFreeSpace(blocks DiskBlocks) DiskBlocks {
	sorted := make(DiskBlocks, len(blocks))
	copy(sorted, blocks)

	head, tail := 0, len(sorted)-1

	for {
		// head went ahead of tail
		if head > tail {
			break
		}

		// if tail pointing to free space we need to move tail forward
		if sorted[tail] == -1 {
			tail--
			continue
		}

		// if head is not pointing to the free space move it ahead
		if sorted[head] != -1 {
			head++
			continue
		}

		// now head should be on a free space and tail on a number
		// switch them
		sorted[head], sorted[tail] = sorted[tail], sorted[head]

		// and move both
		head++
		tail--
	}

	return sorted
}

func MoveFreeSpaceV2(blocks DiskBlocks, freeSpaceHeaps [10]utils.IntHeap) DiskBlocks {
	sorted := make(DiskBlocks, len(blocks))
	copy(sorted, blocks)

	tail := len(sorted) - 1

	for tail >= 0 {
		// if it's free space we basically skip as we do not care about it
		if sorted[tail] == -1 {
			tail--
			continue
		}

		// ID and a size of the file
		idx, size := sorted[tail], 0

		// find the size of a block, loop until we get to the next idx
		for tail >= 0 && sorted[tail] == idx {
			size++
			tail--
		}

		// we know the size of a file and it's ID, we need to find a place for that in the
		// left most free space (begins with a smaller index) where this file fits (free space >= size)
		heapSizeThatFits := -1
		smallestIndex := math.MaxInt

		for heapSize := size; heapSize < 10; heapSize++ {
			usedHeap := freeSpaceHeaps[heapSize]
			if len(usedHeap) == 0 {
				continue
			}

			if usedHeap[0] < smallestIndex {
				heapSizeThatFits = heapSize
				smallestIndex = usedHeap[0]
			}
		}

		// nothing found, so we do not do anything with this number
		if heapSizeThatFits == -1 {
			continue
		}

		// smallest index still need to be smaller than we have as tail
		if smallestIndex > tail {
			continue
		}

		// remove free space from the heap
		heap.Pop(&freeSpaceHeaps[heapSizeThatFits])

		// replace a free space with the number and other way around
		for j := 0; j < size; j++ {
			sorted[j+smallestIndex] = idx
			sorted[j+tail+1] = -1
		}

		// add the rest of a free space to the heaps (of a different size)
		// a new starting point is the current starting point + file size
		// todo: funny bug that does not break anything, `newHeapSize` = 0, then we should not push anywhere, but it still works ;)
		newHeapSize := heapSizeThatFits - size
		heap.Push(&freeSpaceHeaps[newHeapSize], smallestIndex+size)

	}

	return sorted
}

func Checksum(blocks DiskBlocks) int {
	var sum int
	for idx, num := range blocks {
		// free space
		if num == -1 {
			continue
		}
		sum += idx * num
	}
	return sum
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	initialBlocks, freeSpaceHeaps := ReadDiskBlocks(exampleInput)
	fmt.Printf("Initial block:\n%s\n%v\n", initialBlocks, freeSpaceHeaps)
	movedBlockes := MoveFreeSpace(initialBlocks)
	fmt.Printf("Moved block:\n%s\n", movedBlockes)
	movedBlockes2 := MoveFreeSpaceV2(initialBlocks, freeSpaceHeaps)
	fmt.Printf("Moved block v2:\n%s\n", movedBlockes2)
	checkSum := Checksum(movedBlockes)
	fmt.Printf("Checksum:\n%d\n", checkSum)
	checkSum2 := Checksum(movedBlockes2)
	fmt.Printf("Checksum v2:\n%d\n", checkSum2)
}
