package main

import (
	"fmt"
)

// Node represents a node in the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert inserts a new node with the given value into the binary tree
func (n *Node) Insert(value int) {
	if n == nil {
		return
	} else if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrderTraversal traverses the binary tree in-order and prints the values
func (n *Node) InOrderTraversal() {
	if n != nil {
		n.Left.InOrderTraversal()
		fmt.Println(n.Value)
		n.Right.InOrderTraversal()
	}
}

// PreOrderTraversal traverses the binary tree in pre-order and prints the values
func (n *Node) PreOrderTraversal() {
	if n != nil {
		fmt.Println(n.Value)
		n.Left.PreOrderTraversal()
		n.Right.PreOrderTraversal()
	}
}

// PostOrderTraversal traverses the binary tree in post-order and prints the values
func (n *Node) PostOrderTraversal() {
	if n != nil {
		n.Left.PostOrderTraversal()
		n.Right.PostOrderTraversal()
		fmt.Println(n.Value)
	}
}

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(13)
	root.Insert(17)

	fmt.Println("In-Order Traversal:")
	root.InOrderTraversal()

	fmt.Println("Pre-Order Traversal:")
	root.PreOrderTraversal()

	fmt.Println("Post-Order Traversal:")
	root.PostOrderTraversal()
}
