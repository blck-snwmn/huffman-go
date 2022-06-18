package huffmango

import "container/heap"

var _ heap.Interface = (*queue)(nil)

type queue []treeNode

// Len implements heap.Interface
func (t *queue) Len() int {
	return len(*t)
}

// Less implements heap.Interface
func (t *queue) Less(i int, j int) bool {
	return (*t)[i].count() < (*t)[j].count()
}

// Swap implements heap.Interface
func (t *queue) Swap(i int, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

// Pop implements heap.Interface
func (t *queue) Pop() any {
	last := (*t)[len(*t)-1]

	*t = (*t)[:len(*t)-1]

	return last
}

// Push implements heap.Interface
func (t *queue) Push(x any) {
	*t = append(*t, x.(treeNode))
}
