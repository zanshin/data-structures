package main

import (
	"fmt"
)

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
// Search

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
}
