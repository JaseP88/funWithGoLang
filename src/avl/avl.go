// need to be able to add
// need to be able to delete
// need to be able to find

package main

import (
	"fmt"
)

// Node ...
type Node struct {
	key, height         int
	parent, left, right *Node
}

func (node *Node) insert(key int) {
	if node == nil {
		return
	}

	switch {
	case node.key == key:
		return

	case key < node.key:
		if node.left == nil {
			node.left = &Node{key: key, parent: node}
			if node.height < 1 {
				node.height = 1
			}
			node.parent.resetParentHeight(node.height)
			node.balance()
			return
		}
		node.left.insert(key)

	case key > node.key:
		if node.right == nil {
			node.right = &Node{key: key, parent: node}
			if node.height < 1 {
				node.height = 1
			}
			node.parent.resetParentHeight(node.height)
			node.balance()
			return
		}
		node.right.insert(key)
	}
}

func (node *Node) inOrder() {
	if node == nil {
		return
	}
	node.left.inOrder()
	fmt.Print(node.key)
	fmt.Print(": ")
	fmt.Println(node.height)
	node.right.inOrder()
}

func (node *Node) resetParentHeight(h int) {
	if node != nil {
		if h+1 > node.height {
			node.height = h + 1
			node.parent.resetParentHeight(node.height)
		}
	}
}

func (node *Node) balance() {
	if node == nil {
		fmt.Println("DEAD")
		fmt.Println()
	}
	if node != nil {
		leftHeight, rightHeight := node.getSubTreeHeight()
		balance := leftHeight - rightHeight

		if balance > 1 || balance < -1 {
			//balance
			fmt.Print(node.key)
			fmt.Print(": ")
			fmt.Println(balance)
		} else {
			fmt.Print(node.key)
			fmt.Print(": ")
			fmt.Println(balance)
			fmt.Println("going UP")
			node.parent.balance()
		}
	}
}

func (node *Node) getSubTreeHeight() (int, int) {
	var leftHeight, rightHeight int

	if node.left == nil {
		leftHeight = -1
	} else {
		leftHeight = node.left.height
	}
	if node.right == nil {
		rightHeight = -1
	} else {
		rightHeight = node.right.height
	}
	return leftHeight, rightHeight
}

func main() {
	tree := &Node{key: 17}
	tree.insert(12)
	tree.insert(22)
	tree.insert(1)
	tree.insert(7)
	tree.insert(3)

	// fmt.Println("HAHHAHAHAHAH")
	// tree.inOrder()
}
