package bst

type IBinarySearch interface {
	Load(tree []int64)
	Insert(v int64)
	Delete(v int64)
	Traversal(t TraversalType) []int64
	DeepestNodes() ([]int64, int)
	Root() int64
	findParent(v int64) (*Node, *Node)
	// findDeep(n *Node, lvl int, maxLevel int, deepNode *Node)
	min(n *Node) (*Node, *Node)
	max(n *Node) (*Node, *Node)
}

type BinarySearch struct {
	root *Node
}

type TraversalType string

const (
	PreOrder  TraversalType = "Pre Order"
	InOrder   TraversalType = "In Order"
	PostOrder TraversalType = "Post Order"
)

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
func (b *BinarySearch) Root() int64 {
	if b.root != nil {
		return b.root.Value
	}
	return 0
}

// DeepestNode implements IBinarySearch
func (b *BinarySearch) DeepestNodes() ([]int64, int) {
	deepNodes := make(map[int][]int64)
	max := 0
	b.findDeep(b.root, 0, &deepNodes)
	for k := range deepNodes {
		if k > max {
			max = k
		}
	}
	return deepNodes[max], max
}

// Load implements IBinarySearch
func (b *BinarySearch) Load(tree []int64) {
	if len(tree) == 0 {
		return
	}
	for _, n := range tree {
		b.Insert(n)
	}
}

// Insert implements IBinarySearch
func (b *BinarySearch) Insert(v int64) {
	if b.root == nil {
		b.root = &Node{
			Value: v,
		}
		return
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
		return
	}
	ref.Right = &Node{Value: v}
}

// Delete implements IBinarySearch
func (b *BinarySearch) Delete(v int64) {
	parent, n := b.findParent(v)

	// nothing to delete
	if n == nil {
		return
	}

	// n is a leaf
	if n.Left == nil {
		if parent.Left == n {
			parent.Left = n.Right
			return
		}
		if parent.Right == n {
			parent.Right = n.Right
			return
		}
	}
	if n.Right == nil {
		if parent.Left == n {
			parent.Left = n.Left
			return
		}
		if parent.Right == n {
			parent.Right = n.Left
			return
		}

	}

	if n.Right != nil {
		pm, m := b.min(n.Right)
		n.Value = m.Value
		pm.Left = nil
		return
	}

	if n.Left != nil {
		pm, m := b.max(n.Left)
		n.Value = m.Value
		pm.Right = nil
		return
	}
}

// findParent implements IBinarySearch returns (parent,node)
func (b *BinarySearch) findParent(v int64) (*Node, *Node) {
	if b.root == nil {
		return nil, nil
	}
	if b.root.Value == v {
		return nil, b.root
	}
	rf := b.root
	var parent *Node
	for rf != nil {
		parent = rf
		if rf.Left != nil && rf.Left.Value == v {
			return rf, rf.Left
		}
		if rf.Right != nil && rf.Right.Value == v {
			return rf, rf.Right
		}
		if v < rf.Value {
			rf = rf.Left
		} else {
			rf = rf.Right
		}
	}
	return parent, nil
}

// min implements IBinarySearch returns (parent, node)
func (b *BinarySearch) min(n *Node) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}
	rf := n
	var parent *Node
	for rf.Left != nil {
		parent = rf
		rf = rf.Left
	}
	return parent, rf
}

// max implements IBinarySearch returns (parent, node)
func (b *BinarySearch) max(n *Node) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}
	rf := n
	var parent *Node
	for rf.Right != nil {
		parent = rf
		rf = rf.Right
	}
	return parent, rf
}

// findDeep implements IBinarySearch
func (b *BinarySearch) findDeep(n *Node, lvl int, nodes *map[int][]int64) {
	if n == nil {
		return
	}

	(*nodes)[lvl] = append((*nodes)[lvl], n.Value)

	if n.Left != nil {
		b.findDeep(n.Left, lvl+1, nodes)
	}

	if n.Right != nil {
		b.findDeep(n.Right, lvl+1, nodes)
	}
}

// Traversal implements IBinarySearch
func (b *BinarySearch) Traversal(t TraversalType) []int64 {
	var out []int64
	if t == InOrder {
		inOrder(b.root, &out)
		return out
	}
	if t == PreOrder {
		preOrder(b.root, &out)
		return out
	}
	if t == PostOrder {
		postOrder(b.root, &out)
		return out
	}
	return out
}

func inOrder(rt *Node, res *[]int64) {
	if rt == nil {
		return
	}
	inOrder(rt.Left, res)
	*res = append(*res, rt.Value)
	inOrder(rt.Right, res)
}

func preOrder(rt *Node, res *[]int64) {
	if rt == nil {
		return
	}
	*res = append(*res, rt.Value)
	preOrder(rt.Left, res)
	preOrder(rt.Right, res)
}

func postOrder(rt *Node, res *[]int64) {
	if rt == nil {
		return
	}
	postOrder(rt.Left, res)
	postOrder(rt.Right, res)
	*res = append(*res, rt.Value)
}
