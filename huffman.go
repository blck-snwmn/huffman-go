package huffmango

import (
	"container/heap"
	"strings"
)

type treeNode interface {
	count() int
}

var _ treeNode = (*node)(nil)

type node struct {
	left, right *treeNode
	c           int
}

// count implements treeNode
func (n *node) count() int {
	return n.c
}

var _ treeNode = (*leaf)(nil)

type leaf struct {
	v rune
	c int
}

// count implements treeNode
func (l *leaf) count() int {
	return l.c
}

type queue []treeNode

// Len implements heap.Interface
func (t *queue) Len() int {
	return len(*t)
}

// Less implements heap.Interface
func (t *queue) Less(i int, j int) bool {
	return (*t)[i].count() > (*t)[j].count()
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

var _ heap.Interface = (*queue)(nil)

type counter struct {
	r     rune
	count int
}

func count(s string) []counter {
	m := make(map[rune]int, len(s))
	orderStore := make([]rune, 0, len(s))

	for _, r := range s {
		if _, ok := m[r]; !ok {
			orderStore = append(orderStore, r)
		}
		m[r]++
	}

	store := make([]counter, 0, len(s))
	for _, r := range orderStore {
		store = append(store, counter{
			r:     r,
			count: m[r],
		})
	}
	return store
}

func buildTree(c []counter) treeNode {
	t := &queue{}
	heap.Init(t)
	for _, v := range c {
		k, v := v.r, v.count
		heap.Push(t, &leaf{
			v: k,
			c: v,
		})
	}
	for t.Len() > 1 {
		// t saved in ascending order of appearance
		hevy := t.Pop().(treeNode)
		low := t.Pop().(treeNode)

		t.Push(&node{
			left:  &low,
			right: &hevy, // Save the less frequent `node` on the right side
			c:     hevy.count() + low.count(),
		})
	}
	return (*t)[0]
}

func createTable(tn treeNode, prefix string, m map[rune]string) map[rune]string {
	switch tn := tn.(type) {
	case *node:
		if tn.left != nil {
			mm := createTable(*tn.left, prefix+"0", m)
			for k, v := range mm {
				m[k] = v
			}
		}
		if tn.right != nil {
			mm := createTable(*tn.right, prefix+"1", m)
			for k, v := range mm {
				m[k] = v
			}
		}
	case *leaf:
		m[tn.v] = prefix
	}
	return m
}

func createDecodeTable(m map[rune]string) map[string]rune {
	r := map[string]rune{}
	for k, v := range m {
		r[v] = k
	}
	return r
}

func EncodeWithDecodeTable(s string) (string, map[string]rune) {
	if len(s) == 0 {
		return s, nil
	}

	m := count(s)
	t := buildTree(m)

	table := createTable(t, "", map[rune]string{})

	var builder strings.Builder
	for _, ss := range s {
		builder.WriteString(table[ss])
	}
	return builder.String(), createDecodeTable(table)
}
func Encode(s string) string {
	result, _ := EncodeWithDecodeTable(s)
	return result
}
