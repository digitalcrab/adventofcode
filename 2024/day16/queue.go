package main

// Queue (see `container/heap`
type Queue []*Step

func (pq Queue) Len() int { return len(pq) }

func (pq Queue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Queue) Push(x interface{}) {
	item := x.(*Step)
	*pq = append(*pq, item)
}

func (pq *Queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
