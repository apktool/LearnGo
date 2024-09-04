package app

import (
	"fmt"
	"strings"
)

type Node struct {
	part     string
	children []*Node
	isWild   bool
}

func (n *Node) String() string {
	return fmt.Sprintf("node{ part=%s, isWild=%t}", n.part, n.isWild)
}

func (n *Node) Insert(parts []string, height int) {
	if len(parts) == height {
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &Node{part: part, isWild: false}
		n.children = append(n.children, child)
	}

	child.Insert(parts, height+1)
}

func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		return n
	}

	children := n.matchChildren(parts[height])

	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *Node) matchChild(part string) *Node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
