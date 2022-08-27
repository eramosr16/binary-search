package bst

type IBinarySearch interface {
	Insert(v int64) *Node
	Delete(v int64) *Node
	ToArray() []int64
	Root() *Node
	findParent(v int64) *Node
}

type BinarySearch struct {
	root *Node
}

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

// NewBinarySearch returns new BinarySearch.
func NewBinarySearch() *BinarySearch {
	return &BinarySearch{}
}

var _ IBinarySearch = (*BinarySearch)(nil)

// Root implements IBinarySearch
func (b *BinarySearch) Root() *Node {
	return b.root
}

// Insert implements IBinarySearch
func (b *BinarySearch) Insert(v int64) *Node {
	if b.root == nil {
		b.root = &Node{
			Value: v,
		}
		return b.root
	}
	r := b.root
	var ref *Node
	for r != nil {
		ref = r
		if v < r.Value {
			r = r.Left
		} else {
			r = r.Right
		}
	}
	if v < ref.Value {
		ref.Left = &Node{Value: v}
		return ref.Left
	}
	ref.Right = &Node{Value: v}
	return ref.Right
}

// Delete implements IBinarySearch
func (b *BinarySearch) Delete(v int64) *Node {
	parent := b.findParent(v)
	if parent == nil {
		return nil
	}
	
	panic("unimplemented")
}

// find implements IBinarySearch
func (b *BinarySearch) findParent(v int64) *Node {
	if b.root == nil {
		return nil
	}
	if b.root.Value == v {
		return nil
	}
	rf := b.root
	var parent *Node
	for rf != nil {
		parent = rf
		if rf.Left != nil && rf.Left.Value == v {
			return rf
		}
		if rf.Right != nil && rf.Right.Value == v {
			return rf
		}
		if rf.Value < v {
			rf = rf.Left
		} else {
			rf = rf.Right
		}
	}
	return parent
}

// ToArray implements IBinarySearch
func (*BinarySearch) ToArray() []int64 {
	panic("unimplemented")
}
