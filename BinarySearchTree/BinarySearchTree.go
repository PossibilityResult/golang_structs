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


func CreateBSTRoot(_data int) BinarySearchTree {
	root := Node{
		data: _data,
		right: nil,
		left: nil, 
	}

	return BinarySearchTree{
		root: &root,
	}
}

func PrintBSTRoot(_BST BinarySearchTree) {
	fmt.Println(_BST.root.data)
}