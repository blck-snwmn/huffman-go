package huffmango

import "container/heap"

type node struct {
	left, right *node
	value       uint
}

type tree []node

// Len implements heap.Interface
func (t *tree) Len() int {
	return len(*t)
}

// Less implements heap.Interface
func (t *tree) Less(i int, j int) bool {
	return (*t)[i].value < (*t)[j].value
}

// Swap implements heap.Interface
func (t *tree) Swap(i int, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

// Pop implements heap.Interface
func (t *tree) Pop() any {
	last := (*t)[len(*t)-1]

	*t = (*t)[:len(*t)-1]

	return last
}

// Push implements heap.Interface
func (t *tree) Push(x any) {
	*t = append(*t, x.(node))
}

var _ heap.Interface = (*tree)(nil)
