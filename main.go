package main

import (
	"leetcode/tree"
)

func main() {
	node := &tree.TNode{
		Val: 5,
		Left: &tree.TNode{
			Val: 4,
			Left: &tree.TNode{
				Val: 11,
				Left: &tree.TNode{
					Val: 7,
				},
				Right: &tree.TNode{
					Val: 2,
				},
			},
		},
		Right: &tree.TNode{
			Val: 8,
			Left: &tree.TNode{
				Val: 13,
			},
			Right: &tree.TNode{
				Val: 4,
				Left: &tree.TNode{
					Val: 5,
				},
				Right: &tree.TNode{
					Val: 1,
				},
			},
		},
	}
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
