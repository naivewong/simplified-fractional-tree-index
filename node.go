package fti

import (
	"strconv"
)

type Node struct {
	key   int
	value string
	child *Node
}

func NewNode(key int, value string) *Node {
	return &Node{key: key, value: value}
}

func (n *Node) Key() int { return n.key }

func (n *Node) Value() string { return n.value }

func (n *Node) Child() *Node { return n.child }

func (n *Node) SetKey(key int) { n.key = key }

func (n *Node) SetValue(value string) { n.value = value }

func (n *Node) String() string { return "(" + strconv.Itoa(n.key) + "," + n.value + ")" }