package main

import (
	"golang_structs/BinarySearchTree"
)

func main() {
	var BST BinarySearchTree.BinarySearchTree
	BST.Insert(10)
	BST.Insert(5)
	BST.PrintBSTRoot()
}