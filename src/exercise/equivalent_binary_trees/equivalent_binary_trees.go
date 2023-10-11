package equivalentbinarytrees

import (
	"golang.org/x/tour/tree"
)

func WalkRecursive(ch chan int, t *tree.Tree) {
	walkRecursive(ch, t)
	close(ch)
}

func walkRecursive(ch chan int, t *tree.Tree) {
	// Walk walks the tree t sending all values from the tree to the channel ch.
	// 1. Implement the Walk function. Read and print 10 values from the
	//		channel. It should be sorted.
	// 2. Test the Walk function.
	if t == nil {
		return
	}
	walkRecursive(ch, t.Left)  // rn 0
	ch <- t.Value              // rn 1
	walkRecursive(ch, t.Right) // rn 1
	// rn 2
}

func WalkIterative(ch chan int, t *tree.Tree) {
	type trace struct {
		t  *tree.Tree
		rn int
	}
	stack := []*trace{{t, 0}}
	for sp := stack[0]; len(stack) > 0; {
		if sp.t == nil || sp.rn == 2 {
			stack = stack[:len(stack)-1]
		} else if sp.rn == 0 {
			sp.rn++
			stack = append(stack, &trace{sp.t.Left, 0})
		} else if sp.rn == 1 {
			sp.rn++
			ch <- sp.t.Value
			stack = append(stack, &trace{sp.t.Right, 0})
		}
		if len(stack) > 0 {
			sp = stack[len(stack)-1]
		}
	}
	close(ch)
}

// Same determines whether the trees t1 and t2 contain same values.
func Same(t1, t2 *tree.Tree) bool {
	// 3. Implement the Same function using Walk to determine whether t1 and t2
	// 		store the same values.
	// 4. Test the Same function.
	return false
}
