package main

import (
	"leetcode/linkedlists"
)

func main() {
	reorderList(&linkedlists.ListNode{
		Val: 1,
		Next: &linkedlists.ListNode{
			Val: 1,
		},
	})
}

func reorderList(head *linkedlists.ListNode) *linkedlists.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	even, odd := head, head.Next
	evenHead, oddHead := even, odd
	var lastEven *linkedlists.ListNode
	index := 0
	for even != nil && odd != nil {
		if index%2 == 0 {
			even.Next, even = odd.Next, odd.Next
			if even != nil {
				lastEven = even
			}
		} else {
			odd.Next, odd = even.Next, odd.Next
		}
		index++
	}
	if lastEven == nil {
		evenHead.Next = oddHead
	} else {
		lastEven.Next = oddHead
	}
	return evenHead
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
