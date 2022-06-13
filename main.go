package main

import (
	//"fmt"
	"golang_structs/BinarySearchTree"
)

func main() {
	var BST BinarySearchTree.BinarySearchTree
	BST.Insert(10)
	BST.Insert(5)
	BST.Insert(2)
	BST.Insert(7)
	BST.InorderPrint(BST.Root)
}