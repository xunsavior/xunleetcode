package main

import (
	"leetcode/tree"
	"log"
)

func main() {
	node := &tree.TNode{
		Val:  1,
		Left: nil,
		Right: &tree.TNode{
			Val: 2,
			Left: &tree.TNode{
				Val: 3,
			},
			Right: nil,
		},
	}
	log.Println(inorderTraversal94(node))
}

func inorderTraversal94(root *tree.TNode) []int {
	if root == nil {
		return nil
	}
	output := []int{}
	stack := []*tree.TNode{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		l, r := top.Left, top.Right
		if l != nil {
			stack = append(stack, l)
		} else {
			output = append(output, top.Val)
			stack = stack[0 : len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].Left = nil
			}
			if r != nil {
				stack = append(stack, r)
			}
		}
	}
	return output
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
