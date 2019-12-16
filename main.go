package main

import (
	"fmt"
	"leetcode/tree"
	"math"
)

func main() {
	val := test(&tree.TNode{
		Val: 2,
		Left: &tree.TNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &tree.TNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	})
	fmt.Println(val)
}

func test(current *tree.TNode) int {
	if current == nil {
		return 0
	}
	max := math.MinInt32
	tree.Helper333(current, &max) //helper333()
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
