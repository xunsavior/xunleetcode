package main

import (
	"github.com/xunsavior/go-basic-utils/linkedlist"
)

func main() {
	reverseKGroup25(&linkedlist.ListNode{
		Val: 1,
		Next: &linkedlist.ListNode{
			Val: 2,
			Next: &linkedlist.ListNode{
				Val: 3,
				Next: &linkedlist.ListNode{
					Val:  4,
					Next: &linkedlist.ListNode{Val: 5},
				},
			},
		},
	}, 2)
}

func reverseKGroup25(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	i := 0
	reverseHead, res := head, head
	var (
		reverseTail, lastTail *linkedlist.ListNode
	)
	isFirstTime := true
	for head != nil {
		i++
		if i == k {
			// perform reverse here
			reverseTail, head = head, head.Next
			reverseTail.Next = nil
			var prev *linkedlist.ListNode
			start := reverseHead
			for reverseHead != nil {
				prev, reverseHead.Next, reverseHead = reverseHead, prev, reverseHead.Next
			}
			if isFirstTime {
				res, isFirstTime = prev, false
			}
			/*
				important: we must ensure that the tail node of last reversed list ought
				to point to
			*/
			if lastTail != nil {
				lastTail.Next = prev
			}
			lastTail, start.Next, reverseHead = start, head, head
			i = 0
		} else {
			head = head.Next
		}
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
