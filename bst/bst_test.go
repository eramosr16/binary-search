package bst

import (
	"testing"
)

func TestInsertOne(t *testing.T) {
	b := NewBinarySearch()
	var v int64 = 10

	b.Insert(v)
	res := b.Traversal(InOrder)
	expect := []int64{10}

	if b.root == nil {
		t.Fatalf("Root is nil after insert")
	}
	if len(res) != len(expect) {
		t.Fatalf("Missing nodes on resulting tree")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("Insert failed")
		}
	}
}

func TestInsertOneLeft(t *testing.T) {
	b := NewBinarySearch()
	var v int64 = 10
	var left int64 = 5
	b.Insert(v)
	b.Insert(left)

	res := b.Traversal(InOrder)
	expect := []int64{5, 10}

	if b.root == nil {
		t.Fatalf("Root is nil after insert")
	}
	if len(res) != len(expect) {
		t.Fatalf("Missing nodes on resulting tree")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("Insert failed")
		}
	}
}

func TestInsertOneRight(t *testing.T) {
	b := NewBinarySearch()
	var v int64 = 10
	var right int64 = 15
	b.Insert(v)
	b.Insert(right)

	res := b.Traversal(InOrder)
	expect := []int64{10, 15}

	if b.root == nil {
		t.Fatalf("Root is nil after insert")
	}
	if len(res) != len(expect) {
		t.Fatalf("Missing nodes on resulting tree")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("Insert failed")
		}
	}
}

func TestDeleteLeaf(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{50, 30, 20, 40, 70, 60, 80})

	b.Delete(20)
	res := b.Traversal(InOrder)
	expect := []int64{30, 40, 50, 60, 70, 80}
	if len(res) != len(expect) {
		t.Fatalf("Element not deleted")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("Delete action failed")
		}
	}
}

func TestDeleteFullNode(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{50, 30, 20, 40, 70, 60, 80})

	b.Delete(50)
	res := b.Traversal(InOrder)
	expect := []int64{20, 30, 40, 60, 70, 80}
	if len(res) != len(expect) {
		t.Fatalf("Element not deleted")
	}
	for i := range res {
		if res[i] != expect[i] {
			t.Fatalf("Delete action failed")
		}
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

func TestFindDeepCase1(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{12, 11, 90, 82, 7, 9})
	ns, d := b.DeepestNodes()
	if ns == nil {
		t.Fatalf("Node should not be nil")
	}
	if len(ns) != 1 {
		t.Fatalf("Expected 1 result got %v", len(ns))
	}
	if ns[0] != 9 {
		t.Fatalf("Expected 9 got %v", ns[0])
	}
	if d != 3 {
		t.Fatalf("Depth should be 3")
	}
}

func TestFindDeepCase2(t *testing.T) {
	b := NewBinarySearch()
	b.Load([]int64{26, 82, 16, 92, 33})
	ns, d := b.DeepestNodes()

	if len(ns) != 2 {
		t.Fatalf("Expected 2 result got %v", len(ns))
	}
	if ns[0] != 33 {
		t.Fatalf("Expected 33 got %v", ns[0])
	}
	if ns[1] != 92 {
		t.Fatalf("Expected 92 got %v", ns[1])
	}
	if d != 2 {
		t.Fatalf("Depth should be 2")
	}
}
