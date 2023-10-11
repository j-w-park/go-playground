package equivalentbinarytrees_test

import (
	"math/rand"
	"sync"
	"testing"

	equivalentbinarytrees "github.com/j-w-park/go-playground/src/exercise/equivalent_binary_trees"
	"golang.org/x/tour/tree"
)

var numCases = 1_000

func TestWalkRecursive1(t *testing.T) {
	type tcase struct {
		in   *tree.Tree
		want []int
	}

	// given
	testcases := make([]*tcase, numCases)
	chs := make([]chan int, len(testcases))
	for i := range testcases {
		k := 1 + rand.Intn(9999)
		want := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]
		for j := range want {
			want[j] = k * (j + 1)
		}
		testcases[i] = &tcase{in: tree.New(k), want: want}
		chs[i] = make(chan int)
	}

	// when
	for i, tc := range testcases {
		go equivalentbinarytrees.WalkRecursive(chs[i], tc.in)
	}

	// then
	outs := make([][]int, len(testcases))
	for i, ch := range chs {
		tc := testcases[i]
		outs[i] = make([]int, 0, len(tc.want)) // []
		for v := range ch {                    // blocking occurs
			outs[i] = append(outs[i], v)
		}
		for j, want := range tc.want {
			if want != outs[i][j] {
				t.Errorf("out: %v, expected %v", outs[i], tc.want)
				t.FailNow()
			}
		}
	}
}

func TestWalkRecursive2(t *testing.T) {
	type tcase struct {
		in   *tree.Tree
		want []int
	}
	type result struct {
		index  int
		values []int
	}

	// given
	testcases := make([]*tcase, numCases)
	chs := make([]chan int, len(testcases))
	for i := range testcases {
		k := 1 + rand.Intn(9999)
		want := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]
		for j := range want {
			want[j] = k * (j + 1)
		}
		testcases[i] = &tcase{in: tree.New(k), want: want}
		chs[i] = make(chan int)
	}

	// when
	for i, tc := range testcases {
		go equivalentbinarytrees.WalkRecursive(chs[i], tc.in)
	}

	// then
	var wg sync.WaitGroup
	wg.Add(len(testcases))
	chres := make(chan result, len(testcases))
	for i, tc := range testcases {
		go func(i int, tc *tcase, ch chan int) {
			defer wg.Done()
			out := make([]int, 0, len(tc.want))
			for v := range ch {
				out = append(out, v)
			}
			chres <- result{i, out}
		}(i, tc, chs[i])
	}

	wg.Wait()
	close(chres)

	for res := range chres {
		tc := testcases[res.index]
		for i, want := range tc.want {
			if want != res.values[i] {
				t.Errorf("out: %v, expected %v", res.values, tc.want)
				t.FailNow()
			}
		}
	}
}
