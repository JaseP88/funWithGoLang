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
			node.rightLeftRotate()
		} else {
			node.parent.balance()
		}
	}
}

func (node *Node) leftRotate() {
	temp := *node
	pivot := *node.right

	if temp.parent == nil {
		pivot.parent = nil
		temp.parent = &pivot
	} else {
		pivot.parent = temp.parent
		temp.parent = &pivot
	}

	if pivot.left == nil {
		temp.right = nil
		pivot.left = &temp
	} else {
		pivotLeft := *pivot.left
		temp.right = &pivotLeft
		pivot.left = &temp
	}

	temp.height = pivot.right.height
	pivot.height = pivot.right.height + 1

	*node =  pivot
}

func (node *Node) rightRotate() {
	temp := *node
	pivot := *node.left

	if temp.parent == nil {
		pivot.parent = nil
		temp.parent = &pivot
	} else {
		pivot.parent = temp.parent
		temp.parent = &pivot
	}

	if pivot.right == nil {
		temp.left = nil
		pivot.right = &temp
	} else {
		pivotRight := *pivot.right
		temp.left = &pivotRight
		pivot.right = &temp
	}

	temp.height = pivot.left.height
	pivot.height = pivot.left.height + 1

	*node = pivot
}

func (node *Node) rightLeftRotate() {
	pivot := *node.right
	temp := *pivot.left

	if temp.right == nil {
		pivot.left = nil
	} else {
		tempRight := *temp.right
		pivot.left = &tempRight
	}

	temp.parent = node
	temp.right = &pivot

	pivot.parent = &temp

	temp.height = pivot.height
	pivot.height = temp.height - 1

	node.right = &temp
	node.leftRotate()
}

func (node *Node) leftRightRotate() {
	pivot :=  *node.left
	temp := *pivot.right

	if temp.left == nil {
		pivot.right = nil
	} else {
		tempLeft := *temp.left
		pivot.right = &tempLeft
	}

	temp.parent = node
	temp.left = &pivot

	pivot.parent = &temp

	temp.height = pivot.height
	pivot.height = temp.height - 1

	node.left = &temp
	node.leftRightRotate()
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
	tree := &Node{key: 50}
	tree.insert(44)
	tree.insert(70)
	tree.insert(30)
	tree.insert(47)
	tree.insert(60)
	tree.insert(80)
	tree.insert(55)
	tree.insert(65)
	tree.insert(75)
	tree.insert(90)
	tree.insert(66)


	// tree.preOrder()
	// tree.postOrder()
	tree.inOrder()
}

/*
left rotation
1-2-3

39-32-45-12-33-49-52

39-32-45-12-33-42-60-40-44-50-70-77
*/

/*
right left rotation
4-8-6

50-44-70-30-47-60-80-55-75-90-53 --> x look at node 60

50-44-70-30-47-60-80-55-65-75-90-66
*/