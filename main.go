package main

import (
	"golang_structs/BinarySearchTree"
)

func main() {
	BST := BinarySearchTree.CreateBSTRoot(10)
	BinarySearchTree.PrintBSTRoot(BST)
}