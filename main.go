package main

import (
	"leetcode/tree"
	"log"
)

//nums "leetcode/nums"

func main() {
	sampleDPDemo()
}

/*
	f(0) = 0
	f(1) = 1
	f(2) = f(1) + f(0)
	...
	f(n) = f(n-1) + f(n-2)
*/

func sampleDPDemo() {
	root := &tree.TNode{
		Val: 1,
		Left: &tree.TNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &tree.TNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	flatten114(root)
}

func flatten114(root *tree.TNode) {
	ll := &tree.TNode{}
	helper114(root, ll)
	log.Printf("%v", *root)
	log.Printf("%d %d %d", ll.Val, ll.Right.Val, ll.Right.Right.Val)
}

func helper114(current *tree.TNode, linkList *tree.TNode) {
	if current == nil {
		return
	}

	linkList.Val = current.Val
	linkList.Right = &tree.TNode{}

	helper114(current.Left, linkList.Right)
	helper114(current.Right, linkList.Right)
}

func addMap(nums []int, s *[][]int) {
	nums = append(nums, 3)
	*s = append(*s, nums)
	nums = append(nums, 4)
	nums = append(nums, 5)
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
