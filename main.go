package main

import (
	"leetcode/linkedlists"
)

func main() {
	getIntersectionNode(&linkedlists.ListNode{
		Val: 1,
		Next: &linkedlists.ListNode{
			Val: 2,
			Next: &linkedlists.ListNode{
				Val: 3,
				Next: &linkedlists.ListNode{
					Val: 4,
					Next: &linkedlists.ListNode{
						Val: 5,
					},
				},
			},
		},
	}, 2, 4)
}

func getIntersectionNode(head *linkedlists.ListNode, m int, n int) *linkedlists.ListNode {
	start, i := 1, 0
	if m > 1 {
		start = m - 1
	}
	var (
		startNode, endNode, reversedStart, prev *linkedlists.ListNode
	)
	tmp := head
	for tmp != nil {
		i++
		if i == start {
			startNode = tmp
		}
		if i == m {
			reversedStart = tmp
		}
		if i == n {
			endNode = tmp.Next
			tmp.Next = nil
		}
		tmp = tmp.Next
	}

	end := reversedStart
	for reversedStart != nil {
		prev, reversedStart.Next, reversedStart = reversedStart, prev, reversedStart.Next
	}

	if m == start {
		head, end.Next = prev, endNode
	} else {
		startNode.Next, end.Next = prev, endNode
	}

	return head
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
