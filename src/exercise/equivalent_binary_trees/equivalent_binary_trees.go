package equivalentbinarytrees

import (
	"golang.org/x/tour/tree"
)

// WalkRecursive walks the tree t sending all values from the tree to the channel ch.
func WalkRecursive(ch chan int, parent, child *tree.Tree) {
	// 1. Implement the Walk function. Read and print 10 values from the
	//		channel. It should be sorted.
	// 2. Test the Walk function.
	if child == nil {
		return
	}
	WalkRecursive(ch, child, child.Left)
	ch <- child.Value
	WalkRecursive(ch, child, child.Right)
	if parent == nil {
		close(ch)
	}
}

// Same determines whether the trees t1 and t2 contain same values.
func Same(t1, t2 *tree.Tree) bool {
	// 3. Implement the Same function using Walk to determine whether t1 and t2
	// 		store the same values.
	// 4. Test the Same function.
	return false
}
