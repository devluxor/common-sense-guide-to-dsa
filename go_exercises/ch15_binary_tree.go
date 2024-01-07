package main

import "fmt"

func main() {
	root := &TreeNode{data: 3, leftChild: nil, rightChild: nil}
	root.insert(3)
	root.insert(2)
	root.insert(4)
	root.insert(1)
	root.insert(900)
	// fmt.Println(root.search(5))
	// fmt.Println(root.greatest())
	root.traversePrint()
}

type TreeNode struct {
	data int
	leftChild *TreeNode
	rightChild *TreeNode
}

func (node *TreeNode) search(d int) any {
	if node == nil || node.data == d {
		return node
	}

	if d < node.data {
		return node.leftChild.search(d)
	} else {
		return node.rightChild.search(d)
	}
}

func (node *TreeNode) insert(d int) {
	if d < node.data {
		if node.leftChild == nil {
			node.leftChild = &TreeNode{data: d, leftChild: nil, rightChild: nil}
			return
		}

		node.leftChild.insert(d)
	} else if d > node.data {
		if node.rightChild == nil {
			node.rightChild = &TreeNode{data: d, leftChild: nil, rightChild: nil}
			return
		}

		node.rightChild.insert(d)
	}
}

func (node *TreeNode) traversePrint() {
	if node == nil {
		return
	}

	node.leftChild.traversePrint()
	fmt.Println(node.data)
	node.rightChild.traversePrint()
}

func (currentNode *TreeNode) greatest() *TreeNode {
	if currentNode.rightChild == nil {
		return currentNode
	} 

	return currentNode.rightChild.greatest()
}

