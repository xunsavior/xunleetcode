package main

import (
	"fmt"
	"leetcode/tree"
	"math"
)

func main() {
	test(&tree.TNode{
		Val: 1,
		Left: &tree.TNode{
			Val: 2,
			Left: &tree.TNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &tree.TNode{
			Val:  3,
			Left: nil,
			Right: &tree.TNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	})
}

func test(current *tree.TNode) int {
	if current == nil {
		return 0
	}
	r, l := test(current.Left)+1, test(current.Right)+1
	max := int(math.Max(float64(l), float64(r)))
	fmt.Println(current.Val, max)
	return max
}

// Fibonacci
func f(n int, cache []int) int {
	if cache[n] != 0 {
		return cache[n]
	}
	if n < 3 {
		return n
	}
	res := f(n-2, cache) + f(n-1, cache)
	cache[n] = res
	return res
}

func fSlow(n int) int {
	if n < 3 {
		return n
	}
	return fSlow(n-2) + fSlow(n-1)
}
