package linkedlists

/*
Company: Microsoft(6), Amazon(5), Facebook(3), Apple(3), Uber(2)

Given a linked list, swap every two adjacent nodes and return its head.
You may NOT modify the values in the list's nodes, only nodes itself may be changed.

Example:
Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

func swapPairs24(head *ListNode) *ListNode {
	var (
		res, last *ListNode
		nextStart = head
	)
	for nextStart != nil {
		if nextStart.Next == nil || nextStart.Next.Next == nil {
			nextStart = nil
		} else {
			nextStart = nextStart.Next.Next
			head.Next.Next = nil
		}
		// reverse two adjacent nodes
		var prev *ListNode
		for head != nil {
			prev, head.Next, head = head, prev, head.Next
		}
		head = nextStart
		if res == nil {
			res, last = prev, prev.Next
		} else {
			last, last.Next = prev.Next, prev
		}
	}
	return res
}
