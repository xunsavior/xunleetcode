package main

import (
	"fmt"
	"leetcode/tree"
)

func main() {
	val := tree.MaximumAverageSubtree1120(&tree.TNode{
		Val:  2,
		Left: nil,
		Right: &tree.TNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	})
	fmt.Println(val)
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
