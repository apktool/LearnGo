package app

import (
	"fmt"
	"strings"
)

type Node struct {
	pattern  string
	part     string
	children []*Node
	isWild   bool
}

func (n *Node) String() string {
	return fmt.Sprintf("node{ part=%s, isWild=%t}", n.part, n.isWild)
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &Node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.Insert(pattern, parts, height+1)
}

func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
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
