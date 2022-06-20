package BinaryHeap

import (
	"testing"
)

// Test parent
func TestParent(t *testing.T) {
	got := parent(1)
	if got != 0 {
		t.Errorf("parent(1) = %d; want 0", got)
	}
}

func TestLeft(t *testing.T) {
	got := left(0)
	if got != 1 {
		t.Errorf("left(0) = %d; want 1", got)
	}
}

func TestRight(t *testing.T) {
	got := right(0)
	if got != 2 {
		t.Errorf("right(0) = %d; want 2", got)
	}
}

func TestInsert(t *testing.T) {
	bh := BinaryHeap{}
	bh.Insert(1)
	got := len(bh.arr) 
	if got != 1  {
		t.Errorf("len(bh.arr) = %d; want 1", got)
	}
	bh.Insert(2)
	got = len(bh.arr) 
	if got != 2  {
		t.Errorf("right(0) = %d; want 2", got)
	}
}





