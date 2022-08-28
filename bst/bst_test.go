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

func TestInOrderTraversal(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{50, 30, 20, 40, 70, 60, 80})
	res := b.Traversal(InOrder)

	expect := []int64{20, 30, 40, 50, 60, 70, 80}
	if len(res) != len(expect) {
		t.Fatalf("Missing nodes on resulting tree")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("In Order failed")
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{50, 30, 20, 40, 70, 60, 80})
	res := b.Traversal(PreOrder)

	expect := []int64{50, 30, 20, 40, 70, 60, 80}
	if len(res) != len(expect) {
		t.Fatalf("Missing nodes on resulting tree")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("In Order failed")
		}
	}
}

func TestPostOrderTraversal(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{50, 30, 20, 40, 70, 60, 80})
	res := b.Traversal(PostOrder)

	expect := []int64{20, 40, 30, 60, 80, 70, 50}
	if len(res) != len(expect) {
		t.Fatalf("Missing nodes on resulting tree")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("In Order failed")
		}
	}
}