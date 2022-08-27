package bst

import "testing"

func TestInsertOne(t *testing.T) {
	b := NewBinarySearch()
	var v int64 = 10

	rs := b.Insert(v)

	if b.root == nil {
		t.Fatalf("Root is nil after insert")
	}
	if b.root.Value != v {
		t.Fatalf("Insert was not done on root on empty tree")
	}
	if rs.Value != v {
		t.Fatalf("Resulting node doesn't contain inserted value")
	}
	if rs != b.root {
		t.Fatalf("Resulting node doesn't match root after insert")
	}
}

func TestInsertOneLeft(t *testing.T) {
	b := NewBinarySearch()
	var v int64 = 10
	var left int64 = 5
	b.Insert(v)

	rs := b.Insert(left)

	if rs.Value != left {
		t.Fatalf("Resulting node doesn't contain inserted value")
	}
	if rs != b.root.Left {
		t.Fatalf("Resulting node is not on the left of the tree")
	}
}

func TestInsertOneRight(t *testing.T) {
	b := NewBinarySearch()
	var v int64 = 10
	var right int64 = 15
	b.Insert(v)

	rs := b.Insert(right)

	if rs.Value != right {
		t.Fatalf("Resulting node doesn't contain inserted value")
	}
	if rs != b.root.Right {
		t.Fatalf("Resulting node is not on the right of the tree")
	}
}
