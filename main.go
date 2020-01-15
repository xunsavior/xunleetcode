package main

import "leetcode/linkedlists"

func main() {
	swapPairs(&linkedlists.ListNode{
		Val: 1,
		Next: &linkedlists.ListNode{
			Val: 2,
			Next: &linkedlists.ListNode{
				Val: 3,
				Next: &linkedlists.ListNode{
					Val:  4,
					Next: &linkedlists.ListNode{Val: 5},
				},
			},
		},
	})
}

func swapPairs(head *linkedlists.ListNode) *linkedlists.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		res, last *linkedlists.ListNode
		nextStart = head
	)

	for nextStart != nil {
		if nextStart.Next == nil || nextStart.Next.Next == nil {
			nextStart = nil
		} else {
			nextStart = nextStart.Next.Next
			head.Next.Next = nil
		}
		var prev *linkedlists.ListNode
		for head != nil {
			prev, head.Next, head = head, prev, head.Next
		}
		if res == nil {
			res, last = prev, prev.Next
		} else {
			last, last.Next = prev.Next, prev
		}
		head = nextStart
	}
	return res
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
