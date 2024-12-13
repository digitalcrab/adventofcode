package main

// Step describes one step in solving the puzzle
type Step struct {
	cost, y, x int
}

// Steps represent simple min-heap, where `cost` is the indicator of what small is
type Steps []Step

func (h Steps) Len() int           { return len(h) }
func (h Steps) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h Steps) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Steps) Push(x any) {
	*h = append(*h, x.(Step))
}

func (h *Steps) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
