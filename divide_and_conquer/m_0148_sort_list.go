package divideandconquer

import "github.com/xunsavior/go-basic-utils/data_structure/linkedlist"

/*
Company: Microsoft(3), Adobe(2), Amazon(2); Google(3), Facebook(2)
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:
Input: 4->2->1->3
Output: 1->2->3->4

Example 2:
Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/
// time: O(n*logn) space: O(1) bottom up solution
func sortList148(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// get the length of the list
	n := 0
	tmp := head
	for tmp != nil {
		n++
		tmp = tmp.Next
	}

	dummy := &linkedlist.ListNode{}
	dummy.Next = head
	var (
		l, r, tail *linkedlist.ListNode
	)
	for i := 1; i < n; i = i * 2 {
		tmp = dummy.Next
		tail = dummy
		for tmp != nil {
			l = tmp
			r = split148(l, i)
			tmp = split148(r, i)
			tail.Next, tail = merge148WithBottomUpSolution(l, r)
		}
	}
	return dummy.Next
}

// Splits the linked list into first n element and the rest
// Returns the head of rest
func split148(head *linkedlist.ListNode, n int) *linkedlist.ListNode {
	for n-1 > 0 && head != nil {
		n--
		head = head.Next
	}
	var restHead *linkedlist.ListNode
	if head != nil {
		restHead = head.Next
		head.Next = nil
	}
	return restHead
}

// Merge two lists and return the head and tail of the merged lists
func merge148WithBottomUpSolution(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) (mergedHead, mergedTail *linkedlist.ListNode) {
	tail := &linkedlist.ListNode{}
	head := tail
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	for tail.Next != nil {
		tail = tail.Next
	}
	mergedHead, mergedTail = head.Next, tail
	return
}

// time: O(n*logn) space: O(logn)
func sortList148TopDownSolution(head *linkedlist.ListNode) *linkedlist.ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			fast = nil
		} else {
			fast, slow = fast.Next.Next, slow.Next
		}
	}
	l1, l2 := head, slow.Next
	slow.Next = nil
	return merge148WithTopDownSolution(sortList148TopDownSolution(l1), sortList148TopDownSolution(l2))
}

func merge148WithTopDownSolution(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	tail := &linkedlist.ListNode{}
	head := tail
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return head.Next
}
