package BinarySearchTree

import (
	"fmt"
)

type Node struct {
	Data  int
	Right *Node
	Left  *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (BST *BinarySearchTree) Insert(_data int) {
	newNode := &Node{
		Data:  _data,
		Right: nil,
		Left:  nil,
	}

	if BST.Root == nil {
		BST.Root = newNode
		return
	}

	currentNode := BST.Root

	for true {
		if _data > currentNode.Data {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}
			currentNode = currentNode.Right
		} else {
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}
			currentNode = currentNode.Left
		}
	}
}

func (BST *BinarySearchTree) InorderPrint(root *Node) {
	if root == nil {
		return
	}

	BST.InorderPrint(root.Left)
	fmt.Print(root.Data, " ")
	BST.InorderPrint(root.Right)
}

func (BST *BinarySearchTree) PrintBSTRoot() {
	fmt.Println(BST.Root.Data)
}
