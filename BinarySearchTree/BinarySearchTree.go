package BinarySearchTree

import (
	"fmt"
)

type Node struct {
	data int
	right *Node
	left *Node
}

type BinarySearchTree struct {
	root *Node
}

func (BST BinarySearchTree) Insert(_data int) {
	newNode := Node{
		data: _data,
		right: nil,
		left: nil, 
	}

	if (BST.root == nil) {
		BST.root = &newNode
		return
	}

	currentNode := BST.root

	for true {
		if _data > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &newNode
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode = &newNode
				return
			}
			currentNode = currentNode.left
		}
	}
}

func PrintBSTRoot(_BST BinarySearchTree) {
	fmt.Println(_BST.root.data)
}