// need to be able to add
// need to be able to delete
// need to be able to find

package main

import (
	"fmt"
)

// Node ...
type Node struct {
	key, value  int
	left, right *Node
}

func (node *Node) insert(key, value int) {
	if node == nil {
		return
	}

	switch {
	case node.key == key:
		return

	case key < node.key:
		if node.left == nil {
			node.left = &Node{key: key, value: value}
			return
		}
		node.left.insert(key, value)

	case key > node.key:
		if node.right == nil {
			node.right = &Node{key: key, value: value}
			return
		}
		node.right.insert(key, value)
	}
}

func (node *Node)inOrder() {
	if node == nil {
		return
	}
	node.left.inOrder()
	fmt.Println(node.key)
	node.right.inOrder()
}

func main() {
	tree := &Node{key: 22, value: 1}
	tree.insert(7, 2)
	tree.insert(12, 3)
	tree.insert(17, 4)
	tree.insert(1, 5)
	tree.insert(3, 6)

	tree.inOrder()
}
