package huffmango

import (
	"container/heap"
	"fmt"
	"math/bits"
	"strings"
)

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
	tt := make(queue, 0, len(c))
	t := &tt
	heap.Init(t)
	for _, v := range c {
		heap.Push(t, &leaf{
			v: v.r,
			c: v.count,
		})
	}
	for t.Len() > 1 {
		// t saved in ascending order of appearance
		low := heap.Pop(t).(treeNode)
		heavy := heap.Pop(t).(treeNode)

		heap.Push(t, &node{
			left:  &low,
			right: &heavy, // Save the less frequent `node` on the right side
			c:     low.count() + heavy.count(),
		})
	}
	return (*t)[0]
}

type huffmanTable map[rune]string

func (ht huffmanTable) String() string {
	var builder strings.Builder
	builder.WriteString("table\n")
	for k, v := range ht {
		builder.WriteString(fmt.Sprintf("\t-{%s:%s}\n",
			string([]rune{k}),
			v,
		))
	}
	return builder.String()
}

func createTable(tn treeNode, depth int, encodedValue uint64, m huffmanTable) {
	depth += 1
	switch tn := tn.(type) {
	case *node:
		encodedValue <<= 1
		if tn.left != nil {
			createTable(*tn.left, depth, encodedValue, m)
		}
		if tn.right != nil {
			encodedValue |= 1
			createTable(*tn.right, depth, encodedValue, m)
		}
	case *leaf:
		if depth <= 1 && encodedValue == 0 {
			m[tn.v] = "0"
			return
		}

		var builder strings.Builder
		// builder.Grow(64 - bits.LeadingZeros64(encodedValue))
		builder.Grow(depth)

		// Reverse to add from the beginning
		rev := bits.Reverse64(encodedValue)

		// Start bits.LeadingZeros64(encodedValue) because first 1 to the end values encodes
		for i := 64 - depth; i < 64; i++ {
			r := '0'
			if rev>>i&1 == 1 {
				r = '1'
			}
			builder.WriteRune(r)
		}
		m[tn.v] = builder.String()
	}
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

	table := make(huffmanTable, len(m))
	createTable(t, -1, 0, table)

	fmt.Println(t)
	fmt.Println(table)

	var builder strings.Builder
	builder.Grow(len(s))
	for _, ss := range s {
		builder.WriteString(table[ss])
	}
	// return builder.String(), nil
	return builder.String(), createDecodeTable(table)
}

func Encode(s string) string {
	result, _ := EncodeWithDecodeTable(s)
	return result
}
