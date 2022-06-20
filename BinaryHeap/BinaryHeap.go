package BinaryHeap

import (
	"math"
)

func parent(_i int) int {
	return (_i - 1) / 2
}

func left(_i int) int {
	return (2 * _i) + 1
}

func right(_i int) int {
	return (2 * _i) + 2
}

type BinaryHeap struct {
	arr []int
}

func (bh *BinaryHeap) swap(_i int, _j int,) {
	var temp int = bh.arr[_i]
	bh.arr[_i] = bh.arr[_j]
	bh.arr[_j] = temp
}

func (bh *BinaryHeap) heapify(_i int) {
	var i int = _i
	var l int = left(i)
	var r int = right(i)
	var smallest int = i

	if l < len(bh.arr) && bh.arr[l] < smallest {
		smallest = l
	}
	if r < len(bh.arr) && bh.arr[r] < smallest {
		smallest = r
	}
	if smallest != i {
		bh.swap(i, smallest)
		bh.heapify(smallest)
	}
}

func (bh *BinaryHeap) Insert(_elt int) {
	var i int = len(bh.arr)
	var p int = parent(i)
	bh.arr = append(bh.arr, _elt)

	for i != 0 && _elt < bh.arr[p] {
		bh.swap(i, p)
		p = parent(i)
	}
}

func (bh *BinaryHeap) Pop(_elt int) int {
	var top int = bh.arr[0]
	bh.arr[0] = int(math.Inf(1))
	bh.heapify(0)
	return top
}

func ( bh *BinaryHeap) Top() int {
	return bh.arr[0]
}