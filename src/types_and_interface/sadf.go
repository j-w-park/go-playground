package example_types_and_interface

import (
	"errors"
	"fmt"
)

// treeVal defines an unexported marker interface that makes it clear
// which types can be assigned to val in treeNode
type treeVal interface {
	isToken()
}

type treeNode struct {
	val    treeVal
	lchild *treeNode
	rchild *treeNode
}

type number int

func (number) isToken() {}

type operator func(int, int) int

func (operator) isToken() {}

func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

func parse(s string) (*treeNode, error) {
	// not important for our example, so return something hard-coded
	return &treeNode{
		val: operators["+"],
		lchild: &treeNode{
			val:    operators["*"],
			lchild: &treeNode{val: number(5)},
			rchild: &treeNode{val: number(10)},
		},
		rchild: &treeNode{val: number(20)},
	}, nil
}

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("invalid expression")
	case number:
		return int(val), nil // we know that t.val is of type number, so return the int value
	case operator:
		// find the values of the left and right children, then call the process()
		// method on operator to return the result of processing their values.
		if left, err := walkTree(t.lchild); err != nil {
			return 0, err
		} else if right, err := walkTree(t.rchild); err != nil {
			return 0, err
		} else {
			return val.process(left, right), nil
		}
	default:
		return 0, errors.New("unknown node type")
	}
}

func TreeVal() {
	if parseTree, err := parse("5*10+20"); err != nil {
		panic(err)
	} else {
		fmt.Println(walkTree(parseTree))
	}
}
