package main

import (
	"eramosr16/binary_search/bst"
	"fmt"
)

func main() {
	b := bst.NewBinarySearch()
	b.Load([]int64{12, 11, 90, 82, 7, 9})
	ns, d := b.DeepestNodes()
	fmt.Printf("Output: deepest, %v; depth: %v \n", ns, d)

	b.Clear()
	b.Load([]int64{26, 82, 16, 92, 33})
	ns, d = b.DeepestNodes()
	fmt.Printf("Output: deepest, %v; depth: %v", ns, d)
}
