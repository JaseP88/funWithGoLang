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

func (node *Node) resetParentHeight(h int) {
	if node != nil {
		if h+1 > node.height {
			node.height = h + 1
			node.parent.resetParentHeight(node.height)
		}
	}
}

func (node *Node) balance() {
	if node != nil {
		leftHeight, rightHeight := node.getSubTreeHeight()
		balance := leftHeight - rightHeight

		if balance > 1 || balance < -1 {
			// do rotations
			// node.leftRotate()
			node.rotateLeft()
		} else {
			node.parent.balance()
		}
	}
}

func (node *Node) rotateLeft() {
	pivot := *node
	temp := *node.right

	if pivot.parent == nil {
		temp.parent = nil
		pivot.parent = &temp
	} else {
		temp.parent = pivot.parent
		pivot.parent = &temp
	}

	if temp.left == nil {
		pivot.right = nil
		temp.left = &pivot
	} else {
		temp2 := *temp.left
		pivot.right = &temp2
		temp.left = &pivot
	}

	pivot.height = temp.right.height
	temp.height = temp.right.height + 1

	*node =  temp
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

func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.left.postOrder()
	node.right.postOrder()
	fmt.Println(node.key)
}

func (node *Node) preOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	node.left.preOrder()
	node.right.preOrder()
}

func main() {
	tree := &Node{key: 39}
	tree.insert(32)
	tree.insert(45)
	tree.insert(12)
	tree.insert(33)
	tree.insert(42)
	tree.insert(60)
	tree.insert(40)
	tree.insert(44)
	tree.insert(50)
	tree.insert(70)
	tree.insert(77)

	// tree.preOrder()
	// tree.postOrder()
	tree.inOrder()
	// fmt.Println(tree)

	// fmt.Print(tree.left)
	// fmt.Print(" ")
	// fmt.Println()
	// fmt.Println(tree.left.parent)

	// fmt.Print(tree.right)
	// fmt.Print(" ")
	// fmt.Println(tree.right.parent)
}

/*
left rotation
1-2-3

39-32-45-12-33-49-52

39-32-45-12-33-42-60-40-44-50-70-77
*/
