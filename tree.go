package huffmango

import (
	"fmt"
	"strings"
)

type treeNode interface {
	count() int
	String() string
	string(indent string) string
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

func (n *node) string(indent string) string {
	var build strings.Builder
	build.WriteString(fmt.Sprintf("node(%d)\n", n.c))
	write := func(n *treeNode, ss string) {
		if n == nil {
			return
		}
		s := fmt.Sprintf("%s|-%s :%s", indent, ss, (*n).string(indent+"\t"))
		build.WriteString(s)
	}
	write(n.left, "l")
	build.WriteRune('\n')
	write(n.right, "r")

	return build.String()
}

func (n *node) String() string {
	return n.string("")
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

func (n *leaf) String() string {
	return fmt.Sprintf("leaf(%s, %d)", string([]rune{n.v}), n.c)
}

func (n *leaf) string(indent string) string {
	return n.String()
}
