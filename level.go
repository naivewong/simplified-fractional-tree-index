package fti

import (
	"sort"
)

type Level struct {
	nodes []*Node
	max   int
}

func NewLevel(max int) *Level { return &Level{nodes: make([]*Node, 0, max), max: max} }

// Assume not full and no duplication.
func (l *Level) Insert(node *Node) { l.nodes = append(l.nodes, node) }

func (l *Level) Merge(src, dst *Level) {
	i := 0
	j := 0
	for i < l.max && j < src.max {
		if l.nodes[i].Key() <= src.nodes[j].Key() {
			dst.nodes = append(dst.nodes, l.nodes[i])
			i++
		} else {
			dst.nodes = append(dst.nodes, src.nodes[j])
			j++
		}
	}
	for i < l.max {
		dst.nodes = append(dst.nodes, l.nodes[i])
		i++
	}
	for j < src.max {
		dst.nodes = append(dst.nodes, src.nodes[j])
		j++
	}
	l.nodes = l.nodes[:0]
	src.nodes = src.nodes[:0]
}

func (l *Level) Search(key int) *Node {
	i := sort.Search(len(l.nodes), func (j int) bool {
		return l.nodes[j].Key() >= key
	})
	if i < len(l.nodes) && l.nodes[i].Key() == key {
		return l.nodes[i]
	} else {
		return nil
	}
}

func (l *Level) Len() int { return len(l.nodes) }

func (l *Level) Empty() bool { return len(l.nodes) == 0 }

func (l *Level) Full() bool { return len(l.nodes) == l.max }

func (l *Level) Flush() { l.nodes = l.nodes[:0] }

func (l *Level) String() string {
	s := "["
	for _, n := range l.nodes {
		s += n.String()
	}
	return s + "]"
}